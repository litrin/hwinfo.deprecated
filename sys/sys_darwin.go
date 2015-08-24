// +build darwin

package sys

import (
	"fmt"
	"github.com/mickep76/hwinfo/common"
)

// GetInfo return information about a systems memory.
func GetInfo() (Info, error) {
	fields := []string{
		"Model Name",
		"Model Identifier",
		"Serial Number",
		"Boot ROM Version",
		"SMC Version",
	}
	fmt.Printf("%v\n\n", fields)

	i := Info{}
	i.Manufacturer = "Apple Inc."

	o, err := common.ExecCmdFields("/usr/sbin/system_profiler", []string{"SPHardwareDataType"}, ":", fields)
	fmt.Printf("%v\n\n", o)
	if err != nil {
		return Info{}, err
	}

	i.Model = o["Model Name"]
	i.ModelVersion = o["Model Identifier"]
	i.SerialNumber = o["Serial Number"]
	i.BootROMVersion = o["Boot ROM Version"]
	i.SMCVersion = o["SMC Version"]

	return i, nil
}
