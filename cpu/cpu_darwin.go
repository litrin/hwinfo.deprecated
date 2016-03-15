// +build darwin

package cpu

import (
	"strconv"
	"strings"
	"time"

	"github.com/mickep76/hwinfo/common"
)

func (e *envelope) Refresh() error {
	e.cache.LastUpdated = time.Now()
	e.cache.FromCache = false

	o, err := common.ExecCmdFields("/usr/sbin/sysctl", []string{"-a"}, ":", []string{
		"machdep.cpu.core_count",
		"hw.physicalcpu_max",
		"hw.logicalcpu_max",
		"machdep.cpu.brand_string",
		"machdep.cpu.features",
	})
	if err != nil {
		return err
	}

	e.data.CoresPerSocket, err = strconv.Atoi(o["machdep.cpu.core_count"])
	if err != nil {
		return err
	}

	e.data.Physical, err = strconv.Atoi(o["hw.physicalcpu_max"])
	if err != nil {
		return err
	}

	e.data.Logical, err = strconv.Atoi(o["hw.logicalcpu_max"])
	if err != nil {
		return err
	}

	e.data.Sockets = e.data.Physical / e.data.CoresPerSocket
	e.data.ThreadsPerCore = e.data.Logical / e.data.Sockets / e.data.CoresPerSocket
	e.data.Model = o["machdep.cpu.brand_string"]
	e.data.Flags = strings.ToLower(o["machdep.cpu.features"])

	return nil
}
