package hwinfo

import (
	"os"
	"strings"

	"github.com/mickep76/hwinfo/cpu"
	"github.com/mickep76/hwinfo/disks"
	"github.com/mickep76/hwinfo/dock2box"
	"github.com/mickep76/hwinfo/interfaces"
	"github.com/mickep76/hwinfo/lvm"
	"github.com/mickep76/hwinfo/memory"
	"github.com/mickep76/hwinfo/mounts"
	"github.com/mickep76/hwinfo/opsys"
	"github.com/mickep76/hwinfo/pci"
	"github.com/mickep76/hwinfo/routes"
	"github.com/mickep76/hwinfo/sysctl"
	"github.com/mickep76/hwinfo/system"
)

type HWInfo interface {
	Update() error
	GetData() data
	GetCache() cache
	GetCPU() cpu.CPU
	GetDisks() disks.Disks
	GetDock2Box() dock2box.Dock2Box
	GetInterfaces() interfaces.Interfaces
	GetLVM() lvm.LVM
	GetMemory() memory.Memory
	GetMounts() mounts.Mounts
	GetOpSys() opsys.OpSys
	GetPCI() pci.PCI
	GetRoutes() routes.Routes
	GetSysctl() sysctl.Sysctl
	GetSystem() system.System
	GetInterfaces() interfaces.Interfaces
}

type hwInfo struct {
	CPU        cpu.CPU
	Disks      disks.Disks
	Dock2Box   dock2box.Dock2Box
	Interfaces interfaces.Interfaces
	LVM        lvm.LVM
	Memory     memory.Memory
	Mounts     mounts.Mounts
	OpSys      opsys.OpSys
	PCI        pci.PCI
	Routes     routes.Routes
	Sysctl     sysctl.Sysctl
	System     system.System
	data       *data
	cache      *cache
}

type data struct {
	Hostname      string          `json:"hostname"`
	ShortHostname string          `json:"short_hostname"`
	CPU           cpu.Data        `json:"cpu"`
	Disks         disks.Data      `json:"disks"`
	Dock2Box      dock2box.Data   `json:"dock2box"`
	Interfaces    interfaces.Data `json:"interfaces"`
	LVM           lvm.Data        `json:"lvm"`
	Memory        memory.Data     `json:"memory"`
	Mounts        mounts.Data     `json:"mounts"`
	OpSys         opsys.Data      `json:"opsys"`
	PCI           pci.Data        `json:"pci"`
	Routes        routes.Data     `json:"routes"`
	Sysctl        sysctl.Data     `json:"sysctl"`
	System        system.Data     `json:"system"`
}

type cache struct {
	CPU        cpu.Cache        `json:"cpu"`
	Disks      disks.Cache      `json:"disks"`
	Dock2Box   dock2box.Cache   `json:"dock2box"`
	Interfaces interfaces.Cache `json:"interfaces"`
	LVM        lvm.Cache        `json:"lvm"`
	Memory     memory.Cache     `json:"memory"`
	Mounts     mounts.Cache     `json:"mounts"`
	OpSys      opsys.Cache      `json:"opsys"`
	PCI        pci.Cache        `json:"pci"`
	Routes     routes.Cache     `json:"routes"`
	Sysctl     sysctl.Cache     `json:"sysctl"`
	System     system.Cache     `json:"system"`
}

func New() HWInfo {
	return &hwInfo{
		CPU:        cpu.New(),
		Disks:      disks.New(),
		Dock2Box:   dock2box.New(),
		Interfaces: interfaces.New(),
		LVM:        lvm.New(),
		Memory:     memory.New(),
		Mounts:     mounts.New(),
		OpSys:      opsys.New(),
		PCI:        pci.New(),
		Routes:     routes.Routes(),
		Sysctl:     sysctl.Sysctl(),
		System:     system.New(),
		data:       &data{},
		cache:      &cache{},
	}
}

func (h *hwInfo) GetCPU() cpu.CPU {
	return h.CPU
}

func (h *hwInfo) GetDisks() disks.Disks {
	return h.CPU
}

func (h *hwInfo) GetDock2Box() dock2box.Dock2Box {
	return h.Dock2Box
}

func (h *hwInfo) GetInterfaces() interfaces.Interfaces {
	return h.Interfaces
}

func (h *hwInfo) GetLVM() lvm.LVM {
	return h.LVM
}

func (h *hwInfo) GetMemory() memory.Memory {
	return h.Memory
}

func (h *hwInfo) GetMounts() mounts.Mounts {
	return h.Mounts
}

func (h *hwInfo) GetOpSys() opsys.OpSys {
	return h.OpSys
}

func (h *hwInfo) GetPCI() opsys.PCI {
	return h.PCI
}

func (h *hwInfo) GetRoutes() routes.Routes {
	return h.Routes
}

func (h *hwInfo) GetSysctl() sysctl.Sysctl {
	return h.Sysctl
}

func (h *hwInfo) GetSystem() system.System {
	return h.System
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

	disks := disks.New()
	if err := disks.Update(); err != nil {
		return err
	}
	h.data.Disks = disks.GetData()
	h.cache.Disks = disks.GetCache()

	dock2box := dock2box.New()
	if err := dock2box.Update(); err != nil {
		return err
	}
	h.data.Dock2Box = dock2box.GetData()
	h.cache.Dock2Box = dock2box.GetCache()

	mounts := mounts.New()
	if err := mounts.Update(); err != nil {
		return err
	}
	h.data.Mounts = mounts.GetData()
	h.cache.Mounts = mounts.GetCache()

	sysctl := sysctl.New()
	if err := sysctl.Update(); err != nil {
		return err
	}
	h.data.Sysctl = sysctl.GetData()
	h.cache.Sysctl = sysctl.GetCache()

	pci := pci.New()
	if err := pci.Update(); err != nil {
		return err
	}
	h.data.PCI = pci.GetData()
	h.cache.PCI = pci.GetCache()

	routes := routes.New()
	if err := routes.Update(); err != nil {
		return err
	}
	h.data.Routes = routes.GetData()
	h.cache.Routes = routes.GetCache()

	lvm := lvm.New()
	if err := lvm.Update(); err != nil {
		return err
	}
	h.data.LVM = lvm.GetData()
	h.cache.LVM = lvm.GetCache()

	return nil
}
