// +build linux

package system

import (
	"github.com/mickep76/hwinfo/common"
)

type system struct {
	Manufacturer   string `json:"manufacturer"`
	Product        string `json:"product"`
	ProductVersion string `json:"product_version"`
	SerialNumber   string `json:"serial_number"`
	BIOSVendor     string `json:"bios_vendor"`
	BIOSDate       string `json:"bios_date"`
	BIOSVersion    string `json:"bios_version"`
}

func (s *system) Get() error {
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

	s.Manufacturer = o["chassis_vendor"]
	s.Product = o["product_name"]
	s.ProductVersion = o["product_version"]
	s.SerialNumber = o["product_serial"]
	s.BIOSVendor = o["bios_vendor"]
	s.BIOSDate = o["bios_date"]
	s.BIOSVersion = o["bios_version"]

	return nil
}
