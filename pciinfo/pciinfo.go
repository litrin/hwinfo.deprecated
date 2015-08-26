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
	BusNum      string `json:"bus_number"`
	DeviceNum   string `json:"device_number"`
	DeviceFunc  string `json:"device_function"`
	Class       string `json:"class"`
	VendorID    string `json:"vendor_id"`
	DeviceID    string `json:"device_id"`
	SubVendorID string `json:"subsys_vendor_id"`
	SubDeviceID string `json:"subsys_device_id"`
	VendorName  string `json:"vendor_name"`
	DeviceName  string `json:"device_name"`
	SubsysName  string `json:"subsys_name,omitempty"`
}

// Info structure for information about a systems memory.
type Info struct {
	PCI []PCI `json:"pci"`
}

// TODO: Cache PCI database as a map[string]string
func getPCIVendor(vendorID string, deviceID string, subsysVendorID string, subsysDeviceID string) (string, string, string, error) {

	vendorID = strings.Replace(vendorID, "0x", "", 1)
	deviceID = strings.Replace(deviceID, "0x", "", 1)
	subsysVendorID = strings.Replace(subsysVendorID, "0x", "", 1)
	subsysDeviceID = strings.Replace(subsysDeviceID, "0x", "", 1)

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
	subsysName := ""
	cols := 2
	for _, line := range strings.Split(string(o), "\n") {
		vals := strings.SplitN(line, " ", cols)
		if len(vals) < 2 || strings.HasPrefix(line, "#") {
			continue
		}

		for i := 0; i < cols; i++ {
			vals[i] = strings.Trim(vals[i], " \t")
		}

		// Does it need \t
		/*
			id := strings.Trim(vals[0], " \t")
			val := strings.Trim(vals[1], " ")
		*/

		if strings.LastIndex(line, "\t") == -1 && vals[0] == vendorID {
			vendor = vals[1]
			continue
		}

		if vendor != "" && strings.LastIndex(line, "\t") == 0 && vals[0] == deviceID {
			device = vals[1]
			cols = 3
			continue
		}

		if vendor != "" && device != "" && strings.LastIndex(line, "\t") == 1 && vals[0] == subsysVendorID && vals[1] == subsysDeviceID {
			subsysName = vals[2]
			break
		}

		if vendor != "" && strings.LastIndex(line, "\t") == -1 {
			break
		}
	}

	return vendor, device, subsysName, nil
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
			filepath.Join(path, "subsystem_vendor"),
			filepath.Join(path, "subsystem_device"),
		})
		if err != nil {
			return Info{}, err
		}

		vendorID := o["vendor"]
		deviceID := o["device"]
		subVendorID := o["subsystem_vendor"]
		subDeviceID := o["subsystem_device"]
		vendor, device, subsysName, err := getPCIVendor(vendorID, deviceID, subVendorID, subDeviceID)

		i.PCI = append(i.PCI, PCI{
			BusNum:      pci[1],
			DeviceNum:   dev[0],
			DeviceFunc:  dev[1],
			Class:       o["class"],
			VendorID:    vendorID,
			DeviceID:    deviceID,
			SubDeviceID: subDeviceID,
			SubVendorID: subVendorID,
			VendorName:  vendor,
			DeviceName:  device,
			SubsysName:  subsysName,
		})
	}

	return i, nil
}
