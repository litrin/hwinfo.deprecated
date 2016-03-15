package hwinfo

import (
	"os"
	"strings"

	"github.com/mickep76/hwinfo/cpu"
	//	"github.com/mickep76/hwinfo/disks"
	//	"github.com/mickep76/hwinfo/dock2box"
	"github.com/mickep76/hwinfo/interfaces"
	//	"github.com/mickep76/hwinfo/lvm"
	"github.com/mickep76/hwinfo/memory"
	//	"github.com/mickep76/hwinfo/mounts"
	"github.com/mickep76/hwinfo/opsys"
	//	"github.com/mickep76/hwinfo/pci"
	//	"github.com/mickep76/hwinfo/routes"
	//	"github.com/mickep76/hwinfo/sysctl"
	"github.com/mickep76/hwinfo/system"
)

type data struct {
	Hostname      string      `json:"hostname"`
	ShortHostname string      `json:"short_hostname"`
	CPU           interface{} `json:"cpu"`
	//	Disks         interface{}           `json:"disks"`
	//	Dock2Box      interface{}     `json:"dock2box"`
	Interfaces interface{} `json:"interfaces"`
	//	LVM           interface{]               `json:"lvm"`
	Memory interface{} `json:"memory"`
	//	Mounts        interface{}         `json:"mounts"`
	OpSys interface{} `json:"opsys"`
	//	PCI           interface{}               `json:"pci"`
	//	Routes        interface{}         `json:"routes"`
	//	Sysctl        interface{}         `json:"sysctl"`
	System interface{} `json:"system"`
}

type cache struct {
	CPU interface{} `json:"cpu"`
	//  Disks         interface{}           `json:"disks"`
	//  Dock2Box      interface{}     `json:"dock2box"`
	Interfaces interface{} `json:"interfaces"`
	//  LVM           interface{}              `json:"lvm"`
	Memory interface{} `json:"memory"`
	//  Mounts        interface{}         `json:"mounts"`
	OpSys interface{} `json:"opsys"`
	//  PCI           interface{}            `json:"pci"`
	//  Routes        interface{}        `json:"routes"`
	//  Sysctl        interface{}         `json:"sysctl"`
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

	interfaces := interfaces.New()
	if err := interfaces.Update(); err != nil {
		return err
	}
	h.data.Interfaces = interfaces.GetData()
	h.cache.Interfaces = interfaces.GetCache()

	opSys := opsys.New()
	if err := opSys.Update(); err != nil {
		return err
	}
	h.data.OpSys = opSys.GetData()
	h.cache.OpSys = opSys.GetCache()

	return nil
}
