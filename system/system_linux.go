// +build linux

package system

import (
	"time"

	"github.com/mickep76/hwinfo/common"
)

type data struct {
	Manufacturer   string `json:"manufacturer"`
	Product        string `json:"product"`
	ProductVersion string `json:"product_version"`
	SerialNumber   string `json:"serial_number"`
	BIOSVendor     string `json:"bios_vendor"`
	BIOSDate       string `json:"bios_date"`
	BIOSVersion    string `json:"bios_version"`
}

func (e *envelope) Refresh() error {
	e.cache.LastUpdated = time.Now()
	e.cache.FromCache = false

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

	e.data.Manufacturer = o["chassis_vendor"]
	e.data.Product = o["product_name"]
	e.data.ProductVersion = o["product_version"]
	e.data.SerialNumber = o["product_serial"]
	e.data.BIOSVendor = o["bios_vendor"]
	e.data.BIOSDate = o["bios_date"]
	e.data.BIOSVersion = o["bios_version"]

	return nil
}
