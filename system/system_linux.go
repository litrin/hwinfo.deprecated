// +build linux

package system

import (
	"github.com/mickep76/hwinfo/common"
)

type data struct {
	manufacturer   string `json:"manufacturer"`
	product        string `json:"product"`
	productVersion string `json:"product_version"`
	serialNumber   string `json:"serial_number"`
	biosVendor     string `json:"bios_vendor"`
	biosDate       string `json:"bios_date"`
	biosVersion    string `json:"bios_version"`
}

func (e *envelope) Refresh() error {
	e.cache.lastUpdated = time.Now()
	e.cache.fromCache = false

	o, err := common.LoadFiles([]string{
		"/sys/devices/virtual/dmi/id/chassis_vendor",
		"/sys/devices/virtual/dmi/id/product_name",
		"/sys/devices/virtual/dmi/id/product_version",
		"/sys/devices/virtual/dmi/id/product_serial",
		"/sys/devices/virtual/dmi/id/bios_vendor",
		"/sys/devices/virtual/dmi/id/bios_date",
		"/sys/devices/virtual/dmi/id/bios_version",
	})
	if err != nil {
		return err
	}

	e.data.manufacturer = o["chassis_vendor"]
	e.data.product = o["product_name"]
	e.data.productVersion = o["product_version"]
	e.data.serialNumber = o["product_serial"]
	e.data.biosVendor = o["bios_vendor"]
	e.data.biosDate = o["bios_date"]
	e.data.biosVersion = o["bios_version"]

	return nil
}
