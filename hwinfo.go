package hwinfo

import (
	"github.com/mickep76/hwinfo/cpu"
	"github.com/mickep76/hwinfo/disk"
	"github.com/mickep76/hwinfo/lvm"
	"github.com/mickep76/hwinfo/meminfo"
	"github.com/mickep76/hwinfo/mounts"
	"github.com/mickep76/hwinfo/netinfo"
	"github.com/mickep76/hwinfo/osinfo"
	"github.com/mickep76/hwinfo/pciinfo"
	"github.com/mickep76/hwinfo/routes"
	"github.com/mickep76/hwinfo/sysctl"
	"github.com/mickep76/hwinfo/sysinfo"
	"os"
)

// Info structure for information a system.
type Info struct {
	Hostname string           `json:"hostname"`
	CPU      *cpu.CPU         `json:"cpu"`
	Memory   *meminfo.Info    `json:"memory"`
	OS       *osinfo.Info     `json:"os"`
	System   *sysinfo.Info    `json:"system"`
	Network  *netinfo.Info    `json:"network"`
	PCI      *pciinfo.Info    `json:"pci,omitempty"`
	Disks    *disk.Disk       `json:"disks"`
	Routes   *[]routes.Route  `json:"routes"`
	Sysctl   *[]sysctl.Sysctl `json:"sysctl"`
	LVM      *lvm.LVM         `json:"lvm"`
	Mounts   *[]mounts.Mount  `json:"mounts"`
}

// GetInfo return information about a system.
func GetInfo() (Info, error) {
	i := Info{}

	host, err := os.Hostname()
	if err != nil {
		return Info{}, err
	}
	i.Hostname = host

	i2, err := cpu.Get()
	if err != nil {
		return Info{}, err
	}
	i.CPU = &i2

	i3, err := meminfo.GetInfo()
	if err != nil {
		return Info{}, err
	}
	i.Memory = &i3

	i4, err := osinfo.GetInfo()
	if err != nil {
		return Info{}, err
	}
	i.OS = &i4

	i5, err := sysinfo.GetInfo()
	if err != nil {
		return Info{}, err
	}
	i.System = &i5

	i6, err := netinfo.GetInfo()
	if err != nil {
		return Info{}, err
	}
	i.Network = &i6

	i7, err := pciinfo.GetInfo()
	if err != nil {
		return Info{}, err
	}
	i.PCI = &i7

	i8, err := disks.Get()
	if err != nil {
		return Info{}, err
	}
	i.Disks = &i8

	i9, err := routes.Get()
	if err != nil {
		return Info{}, err
	}
	i.Routes = &i9

	i10, err := sysctl.Get()
	if err != nil {
		return Info{}, err
	}
	i.Sysctl = &i10

	i11, err := lvm.Get()
	if err != nil {
		return Info{}, err
	}
	i.LVM = &i11

	i12, err := mounts.Get()
	if err != nil {
		return Info{}, err
	}
	i.Mounts = &i12

	return i, nil
}
