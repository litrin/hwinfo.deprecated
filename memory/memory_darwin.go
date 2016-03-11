// +build darwin

package memory

import (
	"strconv"
	"time"

	"github.com/mickep76/hwinfo/common"
)

type memory struct {
	TotalKB int       `json:"total_kb"`
	TotalGB int       `json:"total_gb"`
	Last    time.Time `json:"last"`
	TTL     int       `json:"ttl_sec"`
	Fresh   bool      `json:"fresh"`
}

// Get memory info.
func (m *memory) get() error {
	o, err := common.ExecCmdFields("/usr/sbin/sysctl", []string{"-a"}, ":", []string{
		"hw.memsize",
	})
	if err != nil {
		return err
	}

	m.TotalKB, err = strconv.Atoi(o["hw.memsize"])
	if err != nil {
		return err
	}
	m.TotalKB = m.TotalKB / 1024
	m.TotalGB = m.TotalKB / 1024 / 1024

	return nil
}
