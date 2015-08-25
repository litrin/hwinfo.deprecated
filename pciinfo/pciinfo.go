package pciinfo

import (
	"fmt"
	"github.com/mickep76/hwinfo/common"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type PCI struct {
	BusNum     string
	DeviceNum  string
	DeviceFunc string
	Class      string
	VendorID   string
	DeviceID   string
	VendorName string
	DeviceName string
}

// Info structure for information about a systems memory.
type Info struct {
	PCI []PCI `json:"pci"`
}

// TODO: Cache PCI database
func getPCIVendor(vendorID string, deviceID string) (string, string, error) {
	fn := "/usr/share/hwdata/pci.ids"
	if _, err := os.Stat(fn); os.IsNotExist(err) {
		return "", "", fmt.Errorf("file doesn't exist: %s", fn)
	}

	o, err := ioutil.ReadFile(fn)
	if err != nil {
		return "", "", err
	}

	vendor := ""
	device := ""
	for _, line := range strings.Split(string(o), "\n") {
		vals := strings.SplitN(line, " ", 2)
		if len(vals) < 2 || strings.HasPrefix(line, "#") {
			continue
		}

		id := strings.Trim(vals[0], " \t")
		val := strings.Trim(vals[1], " ")

		if strings.LastIndex(line, "\t") == -1 && id == vendorID {
			vendor = val
			continue
		}

		if vendor != "" && strings.LastIndex(line, "\t") == 0 && id == deviceID {
			device = val
			break
		}

		if vendor != "" && strings.LastIndex(line, "\t") == -1 {
			break
		}
	}

	return vendor, device, nil
}

// GetInfo return information about PCI devices.
func GetInfo() (Info, error) {
	i := Info{}

	files, err := filepath.Glob("/sys/bus/pci/devices/*")
	if err != nil {
		return Info{}, err
	}

	for _, path := range files {
		pci := strings.Split(path, ":")
		dev := strings.Split(pci[2], ".")

		o, err := common.LoadFiles([]string{
			filepath.Join(path, "class"),
			filepath.Join(path, "vendor"),
			filepath.Join(path, "device"),
		})
		if err != nil {
			return Info{}, err
		}

		vendorID := strings.TrimLeft(o["vendor"], "0x")
		deviceID := strings.TrimLeft(o["device"], "0x")
		vendor, device, err := getPCIVendor(vendorID, deviceID)

		i.PCI = append(i.PCI, PCI{
			BusNum:     pci[1],
			DeviceNum:  dev[0],
			DeviceFunc: dev[1],
			Class:      o["class"],
			VendorID:   vendorID,
			DeviceID:   deviceID,
			VendorName: vendor,
			DeviceName: device,
		})
	}

	return i, nil
}
