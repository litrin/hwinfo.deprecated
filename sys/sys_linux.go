// +build linux

package sys

import (
	"github.com/mickep76/hwinfo/common"
)

// GetInfo return information about a systems memory.
func GetInfo() (Info, error) {
	i := Info{}

	o, err := common.LoadFile(files)
	if err != nil {
		return Info{}, err
	}

	i.Manufacturer = o["/sys/devices/virtual/dmi/id/chassis_vendor"]
	i.Model = o["/sys/devices/virtual/dmi/id/product_name"]
	i.ModelVersion = o["/sys/devices/virtual/dmi/id/product_version"]
	//	i.SerialNumber = o["/sys/devices/virtual/dmi/id/product_serial"]
	i.BIOSVendor = o["/sys/devices/virtual/dmi/id/bios_vendor"]
	i.BIOSDate = o["/sys/devices/virtual/dmi/id/bios_date"]
	i.BIOSVersion = o["/sys/devices/virtual/dmi/id/bios_version"]

	return i, nil
}
