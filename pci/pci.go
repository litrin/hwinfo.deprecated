package pci

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/mickep76/hwinfo/common"
)

type PCI interface {
	SetTTL(int)
	Get() error
	Refresh() error
}

type device struct {
	Slot      string `json:"slot"`
	ClassID   string `json:"class_id"`
	Class     string `json:"class"`
	VendorID  string `json:"vendor_id"`
	DeviceID  string `json:"device_id"`
	Vendor    string `json:"vendor"`
	Device    string `json:"device"`
	SVendorID string `json:"svendor_id"`
	SDeviceID string `json:"sdevice_id"`
	SName     string `json:"sname"`
}

type pci struct {
	Devices []device  `json:"devices"`
	Last    time.Time `json:"last"`
	TTL     int       `json:"ttl_sec"`
	Fresh   bool      `json:"fresh"`
}

// New constructor.
func New() *pci {
	return &pci{
		TTL: 12 * 60 * 60,
	}
}

// Get info.
func (p *pci) Get() error {
	if p.Last.IsZero() {
		if err := p.Refresh(); err != nil {
			return err
		}
	} else {
		expire := p.Last.Add(time.Duration(p.TTL) * time.Second)
		if expire.Before(time.Now()) {
			if err := p.Refresh(); err != nil {
				return err
			}
		} else {
			p.Fresh = false
		}
	}

	return nil
}

// Refresh cache.
func (p *pci) Refresh() error {
	if err := p.get(); err != nil {
		return err
	}
	p.Last = time.Now()
	p.Fresh = true

	return nil
}

// TODO: Cache PCI database as a map[string]string
func getPCIVendor(vendorID string, deviceID string, sVendorID string, sDeviceID string) (string, string, string, error) {
	fn := "/usr/share/hwdata/pci.ids"
	if _, err := os.Stat(fn); os.IsNotExist(err) {
		return "", "", "", fmt.Errorf("file doesn't exist: %s", fn)
	}

	o, err := ioutil.ReadFile(fn)
	if err != nil {
		return "", "", "", err
	}

	vendor := ""
	device := ""
	sName := ""
	cols := 2
	for _, line := range strings.Split(string(o), "\n") {
		vals := strings.SplitN(line, " ", cols)
		if len(vals) < 2 || strings.HasPrefix(line, "#") {
			continue
		}

		for i := 0; i < cols; i++ {
			vals[i] = strings.Trim(vals[i], " \t")
		}

		if strings.LastIndex(line, "\t") == -1 && vals[0] == vendorID {
			vendor = vals[1]
			continue
		}

		if vendor != "" && strings.LastIndex(line, "\t") == 0 && vals[0] == deviceID {
			device = vals[1]
			cols = 3
			continue
		}

		if vendor != "" && device != "" && strings.LastIndex(line, "\t") == 1 && vals[0] == sVendorID && vals[1] == sDeviceID {
			sName = vals[2]
			break
		}

		if vendor != "" && strings.LastIndex(line, "\t") == -1 {
			break
		}
	}

	return vendor, device, sName, nil
}

// TODO: Cache PCI database as a map[string]string
func getPCIClass(classID string) (string, string, string, error) {
	if len(classID) < 6 {
		return "", "", "", fmt.Errorf("Class string is to short: %s", classID)
	}

	subClassID := classID[2:4]
	progIntfID := classID[4:6]
	classID = classID[0:2]

	fn := "/usr/share/hwdata/pci.ids"
	if _, err := os.Stat(fn); os.IsNotExist(err) {
		return "", "", "", fmt.Errorf("file doesn't exist: %s", fn)
	}

	o, err := ioutil.ReadFile(fn)
	if err != nil {
		return "", "", "", err
	}

	class := ""
	subClass := ""
	progIntf := ""
	begins := false
	for _, line := range strings.Split(string(o), "\n") {
		if !begins && strings.HasPrefix(line, "C 01  Mass storage controller") {
			begins = true
		} else if !begins || strings.HasPrefix(line, "#") {
			continue
		}

		vals := strings.SplitN(strings.Replace(line, "C ", "", 1), " ", 2)
		if len(vals) < 2 {
			continue
		}

		for i := 0; i < 2; i++ {
			vals[i] = strings.Trim(vals[i], " \t")
		}

		if strings.LastIndex(line, "\t") == -1 && vals[0] == classID {
			class = vals[1]
			continue
		}

		if class != "" && strings.LastIndex(line, "\t") == 0 && vals[0] == subClassID {
			subClass = vals[1]
			continue
		}

		if class != "" && subClass != "" && strings.LastIndex(line, "\t") == 1 && vals[0] == progIntfID {
			progIntf = vals[1]
			break
		}

		if class != "" && strings.LastIndex(line, "\t") == -1 {
			break
		}
	}

	return class, subClass, progIntf, nil
}

func (p *pci) get() error {
	files, err := filepath.Glob("/sys/bus/pci/devices/*")
	if err != nil {
		return err
	}

	for _, path := range files {
		slot := strings.SplitN(path, ":", 2)

		o, err := common.LoadFiles([]string{
			filepath.Join(path, "class"),
			filepath.Join(path, "vendor"),
			filepath.Join(path, "device"),
			filepath.Join(path, "subsystem_vendor"),
			filepath.Join(path, "subsystem_device"),
		})
		if err != nil {
			return err
		}

		classID := o["class"][2:]
		class, subClass, _, err := getPCIClass(classID)
		if err != nil {
			return err
		}
		if subClass != "" {
			class = subClass
		}

		vendorID := o["vendor"][2:]
		deviceID := o["device"][2:]
		sVendorID := o["subsystem_vendor"][2:]
		sDeviceID := o["subsystem_device"][2:]
		vendor, dev, sName, err := getPCIVendor(vendorID, deviceID, sVendorID, sDeviceID)
		if err != nil {
			return err
		}

		p.Devices = append(p.Devices, device{
			Slot:      slot[1],
			ClassID:   classID,
			Class:     class,
			VendorID:  vendorID,
			DeviceID:  deviceID,
			SDeviceID: sDeviceID,
			SVendorID: sVendorID,
			Vendor:    vendor,
			Device:    dev,
			SName:     sName,
		})
	}

	return nil
}
