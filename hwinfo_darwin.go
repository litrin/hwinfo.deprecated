package hwinfo

import (
	"os"
	"strings"

	//"github.com/mickep76/hwinfo/cpu"
	//	"github.com/mickep76/hwinfo/interfaces"
	//	"github.com/mickep76/hwinfo/memory"
	//	"github.com/mickep76/hwinfo/opsys"
	"github.com/mickep76/hwinfo/system"
)

type data struct {
	Hostname      string `json:"hostname"`
	ShortHostname string `json:"short_hostname"`
	//	CPU           interface{} `json:"cpu"`
	System interface{} `json:"system"`

	//	Memory        memory.Memory         `json:"memory"`
	//	OpSys         opsys.OpSys           `json:"opsys"`
	//	Interfaces    interfaces.Interfaces `json:"interfaces"`
}

type cache struct {
	//	CPU    interface{} `json:"cpu"`
	System interface{} `json:"system"`

	//	Memory     memory.Cached     `json:"memory"`
	//	OpSys      opsys.Cached      `json:"opsys"`
	//	Interfaces interfaces.Cached `json:"interfaces"`
}

func (h *hwInfo) Update() error {
	host, err := os.Hostname()
	if err != nil {
		return err
	}
	h.data.Hostname = host
	h.data.ShortHostname = strings.Split(host, ".")[0]

	//	cpu := cpu.New()
	//	if err := cpu.Update(); err != nil {
	//		return err
	//	}
	//	e.Data.CPU = cpu.GetData()
	//	e.Cache.CPU = cpu.GetCache()

	system := system.New()
	if err := system.Update(); err != nil {
		return err
	}
	h.data.System = system.GetData()
	h.cache.System = system.GetCache()

	return nil
}
