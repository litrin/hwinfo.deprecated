
# hwinfo
    import "github.com/mickep76/hwinfo"







## type Info
``` go
type Info struct {
    CPU *cpu.Info `json:"cpu"`
    Mem *mem.Info `json:"mem"`
    OS  *os.Info  `json:"os"`
    Sys *sys.Info `json:"sys"`
}
```
Info structure for information a system.









### func GetInfo
``` go
func GetInfo() (Info, error)
```
GetInfo return information about a system.










- - -

# cpu
    import "github.com/mickep76/hwinfo/cpu"







## type Info
``` go
type Info struct {
    Model          string `json:"model"`
    Flags          string `json:"flags"`
    Logical        int    `json:"logical"`
    Physical       int    `json:"physical"`
    Sockets        int    `json:"sockets"`
    CoresPerSocket int    `json:"cores_per_socket"`
    ThreadsPerCore int    `json:"threads_per_core"`
}
```
Info structure for information about a systems CPU(s).









### func GetInfo
``` go
func GetInfo() (Info, error)
```
GetInfo return information about a systems CPU(s).










- - -

# mem
    import "github.com/mickep76/hwinfo/mem"







## type Info
``` go
type Info struct {
    TotalKB int `json:"total_kb"`
}
```
Info structure for information about a systems memory.









### func GetInfo
``` go
func GetInfo() (Info, error)
```
GetInfo return information about a systems memory.










- - -

# os
    import "github.com/mickep76/hwinfo/os"







## type Info
``` go
type Info struct {
    Product string `json:"product"`
    Version string `json:"version"`
}
```
Info structure for information about the operating system.









### func GetInfo
``` go
func GetInfo() (Info, error)
```
GetInfo return information about the operating system.










- - -

# sys
    import "github.com/mickep76/hwinfo/sys"







## type Info
``` go
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
```
Info structure for information about a system.









### func GetInfo
``` go
func GetInfo() (Info, error)
```
GetInfo return information about a systems memory.










- - -
