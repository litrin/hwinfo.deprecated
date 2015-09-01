
# hwinfo
    import "github.com/mickep76/hwinfo"







## type Info
``` go
type Info struct {
    Hostname string           `json:"hostname"`
    CPU      *cpuinfo.Info    `json:"cpu"`
    Memory   *meminfo.Info    `json:"memory"`
    OS       *osinfo.Info     `json:"os"`
    System   *sysinfo.Info    `json:"system"`
    Network  *netinfo.Info    `json:"network"`
    PCI      *pciinfo.Info    `json:"pci,omitempty"`
    Disk     *diskinfo.Info   `json:"disk"`
    Routes   *[]routes.Route  `json:"routes"`
    Sysctl   *[]sysctl.Sysctl `json:"sysctl"`
    LVM      *lvm.LVM         `json:"lvm"`
    Mounts   *[]mounts.Mount  `json:"mounts"`
}
```
Info structure for information a system.









### func GetInfo
``` go
func GetInfo() (Info, error)
```
GetInfo return information about a system.










- - -

# common
    import "github.com/mickep76/hwinfo/common"






## func ExecCmd
``` go
func ExecCmd(cmd string, args []string) (string, error)
```
ExecCmd returns output.


## func ExecCmdFields
``` go
func ExecCmdFields(cmd string, args []string, del string, fields []string) (map[string]string, error)
```
ExecCmdFields returns fields from output.


## func LoadFileFields
``` go
func LoadFileFields(fn string, del string, fields []string) (map[string]string, error)
```
LoadFileFields returns fields from file.


## func LoadFiles
``` go
func LoadFiles(files []string) (map[string]string, error)
```
LoadFiles returns field from multiple files.









- - -

# cpuinfo
    import "github.com/mickep76/hwinfo/cpuinfo"







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

# diskinfo
    import "github.com/mickep76/hwinfo/diskinfo"







## type Disk
``` go
type Disk struct {
    Device string `json:"device"`
    Name   string `json:"name"`
    //	Major  int    `json:"major"`
    //	Minor  int    `json:"minor"`
    //	Blocks int    `json:"blocks"`
    SizeGB int `json:"size_gb"`
}
```










## type Info
``` go
type Info struct {
    Disks []Disk `json:"disk"`
}
```
Info structure for information about a systems memory.









### func GetInfo
``` go
func GetInfo() (Info, error)
```









- - -

# lvm
    import "github.com/mickep76/hwinfo/lvm"






## func GetLogVols
``` go
func GetLogVols() ([]LogVol, error)
```

## func GetPhysVols
``` go
func GetPhysVols() ([]PhysVol, error)
```

## func GetVolGrps
``` go
func GetVolGrps() ([]VolGrp, error)
```


## type LVM
``` go
type LVM struct {
    PhysVols *[]PhysVol `json:"phys_vols"`
    LogVols  *[]LogVol  `json:"log_vols"`
    VolGrps  *[]VolGrp  `json:"vol_grps"`
}
```








### func Get
``` go
func Get() (LVM, error)
```



## type LogVol
``` go
type LogVol struct {
    Name   string `json:"name"`
    VolGrp string `json:"vol_grp"`
    Attr   string `json:"attr"`
    SizeGB int    `json:"size_gb"`
}
```










## type PhysVol
``` go
type PhysVol struct {
    Name   string `json:"name"`
    VolGrp string `json:"vol_group"`
    Format string `json:"format"`
    Attr   string `json:"attr"`
    SizeGB int    `json:"size_gb"`
    FreeGB int    `json:"free_gb"`
}
```










## type VolGrp
``` go
type VolGrp struct {
    Name   string `json:"name"`
    Attr   string `json:"attr"`
    SizeGB int    `json:"size_gb"`
    FreeGB int    `json:"free_gb"`
}
```
















- - -

# meminfo
    import "github.com/mickep76/hwinfo/meminfo"







## type Info
``` go
type Info struct {
    TotalGB int `json:"total_gb"`
}
```
Info structure for information about a systems memory.









### func GetInfo
``` go
func GetInfo() (Info, error)
```
GetInfo return information about a systems memory.










- - -

# mounts
    import "github.com/mickep76/hwinfo/mounts"






## func Get
``` go
func Get() ([]Mount, error)
```


## type Mount
``` go
type Mount struct {
    Source  string `json:"source"`
    Target  string `json:"target"`
    FSType  string `json:"fs_type"`
    Options string `json:"options"`
}
```
















- - -

# netinfo
    import "github.com/mickep76/hwinfo/netinfo"







## type Info
``` go
type Info struct {
    Interfaces    []Interface `json:"interfaces"`
    OnloadVersion string      `json:"onload_version,omitempty"`
}
```
Info structure for information about a systems network.









### func GetInfo
``` go
func GetInfo() (Info, error)
```
GetInfo return information about a systems memory.




## type Interface
``` go
type Interface struct {
    Name            string   `json:"name"`
    MTU             int      `json:"mtu"`
    IPAddr          []string `json:"ipaddr"`
    HWAddr          string   `json:"hwaddr"`
    Flags           []string `json:"flags"`
    Driver          string   `json:"driver,omitempty"`
    DriverVersion   string   `json:"driver_version,omitempty"`
    FirmwareVersion string   `json:"firmware_version,omitempty"`
    PCIBus          string   `json:"pci_bus,omitempty"`
    PCIBusURL       string   `json:"pci_bus_url,omitempty"`
    SwChassisID     string   `json:"sw_chassis_id"`
    SwName          string   `json:"sw_name"`
    SwDescr         string   `json:"sw_descr"`
    SwPortID        string   `json:"sw_port_id"`
    SwPortDescr     string   `json:"sw_port_descr"`
    SwVLAN          string   `json:"sw_vlan"`
}
```
Info structure for information about a systems network interfaces.

















- - -

# osinfo
    import "github.com/mickep76/hwinfo/osinfo"







## type Info
``` go
type Info struct {
    Kernel         string `json:"kernel"`
    KernelVersion  string `json:"kernel_version"`
    Product        string `json:"product"`
    ProductVersion string `json:"product_version"`
}
```
Info structure for information about the operating system.









### func GetInfo
``` go
func GetInfo() (Info, error)
```
GetInfo return information about the operating system.










- - -

# pciinfo
    import "github.com/mickep76/hwinfo/pciinfo"







## type Info
``` go
type Info struct {
    PCI []PCI `json:"pci"`
}
```
Info structure for information about a systems memory.









### func GetInfo
``` go
func GetInfo() (Info, error)
```
GetInfo return information about PCI devices.




## type PCI
``` go
type PCI struct {
    Slot      string `json:"slot"`
    ClassID   string `json:"class_id"`
    Class     string `json:"class"`
    VendorID  string `json:"vendor_id"`
    DeviceID  string `json:"device_id"`
    Vendor    string `json:"vendor"`
    Device    string `json:"device"`
    SVendorID string `json:"svendor_id"`
    SDeviceID string `json:"sdevice_id"`
    SName     string `json:"sname,omiempty"`
}
```
















- - -

# routes
    import "github.com/mickep76/hwinfo/routes"






## func Get
``` go
func Get() ([]Route, error)
```


## type Route
``` go
type Route struct {
    Destination string `json:"destination"`
    Gateway     string `json:"gateway"`
    Genmask     string `json:"genmask"`
    Flags       string `json:"flags"`
    MSS         int    `json:"mss"` // Maximum segment size
    Window      int    `json:"window"`
    IRTT        int    `json:"irtt"` // Initial round trip time
    Interface   string `json:"interface"`
}
```
Info structure for system routes.

















- - -

# run








- - -

# sysctl
    import "github.com/mickep76/hwinfo/sysctl"






## func Get
``` go
func Get() ([]Sysctl, error)
```


## type Sysctl
``` go
type Sysctl struct {
    Key   string `json:"key"`
    Value string `json:"value"`
}
```
Sysctl structure for sysctl key/values.

















- - -

# sysinfo
    import "github.com/mickep76/hwinfo/sysinfo"







## type Info
``` go
type Info struct {
    Manufacturer   string `json:"manufacturer"`
    Product        string `json:"product"`
    ProductVersion string `json:"product_version"`
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
