// +build darwin

package memory

import (
	"github.com/mickep76/hwinfo/common"
	"strconv"
)

// Get information about system memory.
func (m *memoryS) get() error {
	o, err := common.ExecCmdFields("/usr/sbin/sysctl", []string{"-a"}, ":", []string{
		"hw.memsize",
	})
	if err != nil {
		return err
	}

	m.TotalGB, err = strconv.Atoi(o["hw.memsize"])
	if err != nil {
		return err
	}
	m.TotalGB = m.TotalGB / 1024 / 1024 / 1024

	return nil
}
