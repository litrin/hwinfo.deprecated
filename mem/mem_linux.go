// +build linux

package mem

import (
	"github.com/mickep76/hwinfo/common"
	"strconv"
)

// GetInfo return information about a systems CPU(s).
func GetInfo() (Info, error) {
	fields := []string{
		"MemTotal",
	}

	m := Info{}

	o, err := common.LoadFileFields("/proc/meminfo", ":", fields)
	if err != nil {
		return Info{}, err
	}

	m.TotalKB, err = strconv.Atoi(o["MemTotal"])
	if err != nil {
		return Info{}, err
	}

	return m, nil
}
