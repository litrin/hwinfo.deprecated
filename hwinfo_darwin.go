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

type Cached interface {
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

type envelope struct {
	data   HWInfo `json:"data"`
	cached Cached `json:"cached"`
}

type cached struct {
	Hostname      string            `json:"hostname"`
	ShortHostname string            `json:"short_hostname"`
	CPU           cpu.Cached        `json:"cpu"`
	Memory        memory.Cached     `json:"memory"`
	OpSys         opsys.Cached      `json:"opsys"`
	System        system.Cached     `json:"system"`
	Interfaces    interfaces.Cached `json:"interfaces"`
}

func New() HWInfo {
	return &hwInfo{}
}

func NewCached() Cached {
	return &cached{}
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

func (c *cached) Get() error {
	host, err := os.Hostname()
	if err != nil {
		return err
	}
	c.Hostname = host
	c.ShortHostname = strings.Split(host, ".")[0]

	c.CPU = cpu.NewCached()
	if err := c.CPU.Get(); err != nil {
		return err
	}

	c.Memory = memory.NewCached()
	if err := c.Memory.Get(); err != nil {
		return err
	}

	c.OpSys = opsys.NewCached()
	if err := c.OpSys.Get(); err != nil {
		return err
	}

	c.System = system.NewCached()
	if err := c.System.Get(); err != nil {
		return err
	}

	c.Interfaces = interfaces.NewCached()
	if err := c.Interfaces.Get(); err != nil {
		return err
	}

	return nil
}
