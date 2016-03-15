// +build darwin

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
	BootROMVersion string `json:"boot_rom_version"`
	SMCVersion     string `json:"smc_version"`
}

func (e *envelope) Refresh() error {
	e.cache.LastUpdated = time.Now()
	e.cache.FromCache = false
	e.data.Manufacturer = "Apple Inc."

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

	e.data.Product = o["Model Name"]
	e.data.ProductVersion = o["Model Identifier"]
	e.data.SerialNumber = o["Serial Number"]
	e.data.BootROMVersion = o["Boot ROM Version"]
	e.data.SMCVersion = o["SMC Version"]

	return nil
}
