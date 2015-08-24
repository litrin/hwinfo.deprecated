package sys

// Info structure for information about a system.
type Info struct {
	Manufacturer   string `json:"manufacturer"`
	Model          string `json:"model"`
	ModelVersion   string `json:"model_version"`
	SerialNumber   string `json:"serial_number"`
	BIOSVendor     string `json:"bios_vendor,omitempty"`
	BIOSDate       string `json:"bios_date,omitempty"`
	BIOSVersion    string `json:"bios_version,omitempty"`
	BootROMVersion string `json:"boot_rom_version,omitempty"`
	SMCVersion     string `json:"smc_version,omitempty"`
}
