// +build linux

package memory

import (
	"github.com/mickep76/hwinfo/common"
	"strconv"
	"strings"
)

// Get memory info.
func (m *memory) get() error {
	o, err := common.LoadFileFields("/proc/meminfo", ":", []string{
		"MemTotal",
		"MemFree",
		"MemAvailable",
		"Cached",
		"Committed_AS",
		"HugePages_Total",
		"HugePages_Free",
		"HugePages_Rsvd",
		"HugePages_Surp",
		"Hugepagesize",
	})
	if err != nil {
		return err
	}

	m.TotalKB, err = strconv.Atoi(strings.TrimRight(o["MemTotal"], " kB"))
	if err != nil {
		return err
	}
	m.TotalGB = m.TotalKB / 1024 / 1024

	m.FreeKB, err = strconv.Atoi(strings.TrimRight(o["MemFree"], " kB"))
	if err != nil {
		return err
	}
	m.FreeGB = m.FreeKB / 1024 / 1024

	m.AvailableKB, err = strconv.Atoi(strings.TrimRight(o["MemAvailable"], " kB"))
	if err != nil {
		return err
	}
	m.AvailableGB = m.AvailableKB / 1024 / 1024

	m.CachedKB, err = strconv.Atoi(strings.TrimRight(o["Cached"], " kB"))
	if err != nil {
		return err
	}
	m.CachedGB = m.CachedKB / 1024 / 1024

	m.CommittedActSizeKB, err = strconv.Atoi(strings.TrimRight(o["Committed_AS"], " kB"))
	if err != nil {
		return err
	}
	m.CommittedActSizeGB = m.CommittedActSizeKB / 1024 / 1024

	m.HugePagesTot, err = strconv.Atoi(o["HugePages_Total"])
	if err != nil {
		return err
	}

	m.HugePagesFree, err = strconv.Atoi(o["HugePages_Free"])
	if err != nil {
		return err
	}

	m.HugePagesRsvd, err = strconv.Atoi(o["HugePages_Rsvd"])
	if err != nil {
		return err
	}

	m.HugePageSizeKB, err = strconv.Atoi(strings.TrimRight(o["Hugepagesize"], " kB"))
	if err != nil {
		return err
	}

	return nil
}
