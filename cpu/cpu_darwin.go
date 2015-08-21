// +build darwin

package cpu

import (
	"common"
	common "github.com/mickep76/hwinfo/common"
	"strings"
)

// Info return information about a systems CPU(s).
func Info() (Info, error) {
	fields := []string{
		"machdep.cpu.core_count",
		"hw.physicalcpu_max",
		"hw.logicalcpu_max",
		"machdep.cpu.brand_string",
		"machdep.cpu.features",
	}

	c := CPUInfo{}

	o, err := common.ExecCmdFields("/usr/sbin/sysctl", []string{"-a"}, ":", fields)
	if err != nil {
		return CPUInfo, err
	}

	c.CoresPerSocket, err = strconv.ParseInt(o["machdep.cpu.core_count"], 10, 0)
	if err != nil {
		return CPUInfo, err
	}

	c.Physical, err = strconv.ParseInt(o["hw.physicalcpu_max"], 10, 0)
	if err != nil {
		return CPUInfo, err
	}

	c.Logical, err = strconv.ParseInt(o["hw.logicalcpu_max"], 10, 0)
	if err != nil {
		return CPUInfo, err
	}

	c.Sockets = c.Physical / c.CoresPerSocket
	c.ThreadsPerCore = c.Logical / c.Sockets / c.CoresPerSocket
	c.Flags = strings.ToLower(o["cpu_flags"])
}
