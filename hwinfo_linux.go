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
	Get() error
}

type hwInfo struct {
	Hostname      string                `json:"hostname"`
	ShortHostname string                `json:"short_hostname"`
	CPU           cpu.CPU               `json:"cpu"`
	Disks         disks.Disks           `json:"disks"`
	Dock2Box      dock2box.Dock2Box     `json:"dock2box"`
	Interfaces    interfaces.Interfaces `json:"interfaces"`
	LVM           lvm.LVM               `json:"lvm"`
	Memory        memory.Memory         `json:"memory"`
	Mounts        mounts.Mounts         `json:"mounts"`
	OpSys         opsys.OpSys           `json:"opsys"`
	PCI           pci.PCI               `json:"pci"`
	Routes        routes.Routes         `json:"routes"`
	Sysctl        sysctl.Sysctl         `json:"sysctl"`
	System        system.System         `json:"system"`
}

func New() HWInfo {
	return &hwInfo{}
}

func (hwi *hwInfo) Get() error {
	host, err := os.Hostname()
	if err != nil {
		return err
	}
	hwi.Hostname = host
	hwi.ShortHostname = strings.Split(host, ".")[0]

	hwi.CPU = cpu.New()
	if err := hwi.CPU.Get(); err != nil {
		return err
	}

	hwi.Disks = disks.New()
	if err := hwi.Disks.Get(); err != nil {
		return err
	}

	hwi.Dock2Box = dock2box.New()
	if err := hwi.Dock2Box.Get(); err != nil {
		return err
	}

	hwi.Interfaces = interfaces.New()
	if err := hwi.Interfaces.Get(); err != nil {
		return err
	}

	hwi.LVM = lvm.New()
	if err := hwi.LVM.Get(); err != nil {
		return err
	}

	hwi.Memory = memory.New()
	if err := hwi.Memory.Get(); err != nil {
		return err
	}

	hwi.Mounts = mounts.New()
	if err := hwi.Mounts.Get(); err != nil {
		return err
	}

	hwi.OpSys = opsys.New()
	if err := hwi.OpSys.Get(); err != nil {
		return err
	}

	hwi.PCI = pci.New()
	if err := hwi.PCI.Get(); err != nil {
		return err
	}

	hwi.Routes = routes.New()
	if err := hwi.Routes.Get(); err != nil {
		return err
	}

	hwi.Sysctl = sysctl.New()
	if err := hwi.Sysctl.Get(); err != nil {
		return err
	}

	hwi.System = system.New()
	if err := hwi.System.Get(); err != nil {
		return err
	}

	return nil
}
