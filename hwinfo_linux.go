package hwinfo

import (
	"os"
	"strings"

	"github.com/mickep76/hwinfo/cpu"
	//	"github.com/mickep76/hwinfo/disks"
	//	"github.com/mickep76/hwinfo/dock2box"
	//	"github.com/mickep76/hwinfo/interfaces"
	//	"github.com/mickep76/hwinfo/lvm"
	"github.com/mickep76/hwinfo/memory"
	//	"github.com/mickep76/hwinfo/mounts"
	//	"github.com/mickep76/hwinfo/opsys"
	//	"github.com/mickep76/hwinfo/pci"
	//	"github.com/mickep76/hwinfo/routes"
	//	"github.com/mickep76/hwinfo/sysctl"
	"github.com/mickep76/hwinfo/system"
)

type data struct {
	Hostname      string      `json:"hostname"`
	ShortHostname string      `json:"short_hostname"`
	CPU           interface{} `json:"cpu"`
	//	Disks         disks.Disks           `json:"disks"`
	//	Dock2Box      dock2box.Dock2Box     `json:"dock2box"`
	//	Interfaces    interfaces.Interfaces `json:"interfaces"`
	//	LVM           lvm.LVM               `json:"lvm"`
	Memory interface{} `json:"memory"`
	//	Mounts        mounts.Mounts         `json:"mounts"`
	//	OpSys         opsys.OpSys           `json:"opsys"`
	//	PCI           pci.PCI               `json:"pci"`
	//	Routes        routes.Routes         `json:"routes"`
	//	Sysctl        sysctl.Sysctl         `json:"sysctl"`
	System interface{} `json:"system"`
}

type cache struct {
	CPU interface{} `json:"cpu"`
	//  Disks         disks.Disks           `json:"disks"`
	//  Dock2Box      dock2box.Dock2Box     `json:"dock2box"`
	//  Interfaces    interfaces.Interfaces `json:"interfaces"`
	//  LVM           lvm.LVM               `json:"lvm"`
	Memory interface{} `json:"memory"`
	//  Mounts        mounts.Mounts         `json:"mounts"`
	//  OpSys         opsys.OpSys           `json:"opsys"`
	//  PCI           pci.PCI               `json:"pci"`
	//  Routes        routes.Routes         `json:"routes"`
	//  Sysctl        sysctl.Sysctl         `json:"sysctl"`
	System interface{} `json:"system"`
}

func (h *hwInfo) Update() error {
	host, err := os.Hostname()
	if err != nil {
		return err
	}
	h.data.Hostname = host
	h.data.ShortHostname = strings.Split(host, ".")[0]

	cpu := cpu.New()
	if err := cpu.Update(); err != nil {
		return err
	}
	h.data.CPU = cpu.GetData()
	h.cache.CPU = cpu.GetCache()

	system := system.New()
	if err := system.Update(); err != nil {
		return err
	}
	h.data.System = system.GetData()
	h.cache.System = system.GetCache()

	memory := memory.New()
	if err := memory.Update(); err != nil {
		return err
	}
	h.data.Memory = memory.GetData()
	h.cache.Memory = memory.GetCache()

	return nil
}
