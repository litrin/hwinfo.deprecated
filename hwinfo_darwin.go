package hwinfo

import (
	"os"
	"strings"

	"github.com/mickep76/hwinfo/cpu"
	"github.com/mickep76/hwinfo/interfaces"
	"github.com/mickep76/hwinfo/memory"
	"github.com/mickep76/hwinfo/opsys"
	"github.com/mickep76/hwinfo/system"
)

type HWInfo interface {
	Get() error
}

type hwInfo struct {
	Hostname      string                `json:"hostname"`
	ShortHostname string                `json:"short_hostname"`
	CPU           cpu.CPU               `json:"cpu"`
	Memory        memory.Memory         `json:"memory"`
	OpSys         opsys.OpSys           `json:"opsys"`
	System        system.System         `json:"system"`
	Interfaces    interfaces.Interfaces `json:"interfaces"`
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

	hwi.Memory = memory.New()
	if err := hwi.Memory.Get(); err != nil {
		return err
	}

	hwi.OpSys = opsys.New()
	if err := hwi.OpSys.Get(); err != nil {
		return err
	}

	hwi.System = system.New()
	if err := hwi.System.Get(); err != nil {
		return err
	}

	hwi.Interfaces = interfaces.New()
	if err := hwi.Interfaces.Get(); err != nil {
		return err
	}

	return nil
}
