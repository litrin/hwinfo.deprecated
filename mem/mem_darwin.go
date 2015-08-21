// +build darwin

package mem

import (
	"github.com/mickep76/hwinfo/common"
	"strconv"
)

// GetInfo return information about a systems CPU(s).
func GetInfo() (Info, error) {
	fields := []string{
		"hw.memsize",
	}

	m := Info{}

	o, err := common.ExecCmdFields("/usr/sbin/sysctl", []string{"-a"}, ":", fields)
	if err != nil {
		return Info{}, err
	}

	m.TotalKB, err = strconv.Atoi(o["hw.memsize"])
	m.TotalKB = m.TotalKB / 1024
	if err != nil {
		return Info{}, err
	}

	return m, nil
}
