// +build darwin

package system

import (
	"github.com/mickep76/hwinfo/common"
)

type system struct {
	Manufacturer   string `json:"manufacturer"`
	Product        string `json:"product"`
	ProductVersion string `json:"product_version"`
	SerialNumber   string `json:"serial_number"`
	BootROMVersion string `json:"boot_rom_version"`
	SMCVersion     string `json:"smc_version"`
}

func (s *system) Get() error {
	s.Manufacturer = "Apple Inc."

	o, err := common.ExecCmdFields("/usr/sbin/system_profiler", []string{"SPHardwareDataType"}, ":", []string{
		"Model Name",
		"Model Identifier",
		"Serial Number",
		"Boot ROM Version",
		"SMC Version",
	})
	if err != nil {
		return err
	}

	s.Product = o["Model Name"]
	s.ProductVersion = o["Model Identifier"]
	s.SerialNumber = o["Serial Number"]
	s.BootROMVersion = o["Boot ROM Version"]
	s.SMCVersion = o["SMC Version"]

	return nil
}
