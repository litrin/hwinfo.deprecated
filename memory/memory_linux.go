// +build linux

package memory

import (
	"strconv"
	"strings"
	"time"

	"github.com/mickep76/hwinfo/common"
)

/*
MemTotal:       65797620 kB
MemFree:        55647440 kB
MemAvailable:   61837964 kB
Buffers:          950804 kB
Cached:          7070508 kB
SwapCached:            0 kB
Active:          3231820 kB
Inactive:        5093844 kB
Active(anon):    1916540 kB
Inactive(anon):  1064024 kB
Active(file):    1315280 kB
Inactive(file):  4029820 kB
Unevictable:       81808 kB
Mlocked:           81808 kB
SwapTotal:       4194300 kB
SwapFree:        4194300 kB
Dirty:                20 kB
Writeback:             0 kB
AnonPages:        385924 kB
Mapped:           100044 kB
Shmem:           2671360 kB
Slab:            1260964 kB
SReclaimable:    1183296 kB
SUnreclaim:        77668 kB
KernelStack:        7680 kB
PageTables:         6752 kB
NFS_Unstable:          0 kB
Bounce:                0 kB
WritebackTmp:          0 kB
CommitLimit:    37080820 kB
Committed_AS:    3979972 kB
VmallocTotal:   34359738367 kB
VmallocUsed:      441840 kB
VmallocChunk:   34325192504 kB
HardwareCorrupted:     0 kB
AnonHugePages:    120832 kB
HugePages_Total:      12
HugePages_Free:       12
HugePages_Rsvd:        0
HugePages_Surp:        0
Hugepagesize:       2048 kB
DirectMap4k:      142448 kB
DirectMap2M:     8235008 kB
DirectMap1G:    60817408 kB
*/

type memory struct {
	TotalKB             int `json:"total_kb"`
	TotalGB             int `json:"total_gb"`
	FreeKB              int `json:"free_kb"`
	FreeGB              int `json:"free_gb"`
	AvailableKB         int `json:"available_kb"`
	AvailableGB         int `json:"available_gb"`
	CachedKB            int `json:"cached_kb"`
	CachedGB            int `json:"cached_gb"`
	SwapCachedKB        int `json:"swap_cached_kb"`
	SwapCachedGB        int `json:"swap_cached_gb"`
	ActiveKB            int `json:"active_kb"`
	ActiveGB            int `json:"active_gb"`
	InactiveKB          int `json:"inactive_kb"`
	InactiveGB          int `json:"inactive_gb"`
	UnevictableKB       int `json:"unevictable_kb"`
	UnevictableGB       int `json:"unevictable_gb"`
	MLockedKB           int `json:"m_locked_kb"`
	MLockedGB           int `json:"m_locked_gb"`
	SwapTotalKB         int `json:"swap_total_kb"`
	SwapTotalGB         int `json:"swap_total_gb"`
	SwapFreeKB          int `json:"swap_free_kb"`
	SwapFreeGB          int `json:"swap_free_gb"`
	DirtyKB             int `json:"dirty_kb"`
	DirtyGB             int `json:"dirty_gb"`
	WritebackKB         int `json:"writeback_kb"`
	WritebackGB         int `json:"writeback_gb"`
	AnonPagesKB         int `json:"anon_pages_kb"`
	AnonPagesGB         int `json:"anon_pages_gb"`
	MappedKB            int `json:"mapped_kb"`
	MappedGB            int `json:"mapped_gb"`
	ShmemKB             int `json:"shmem_kb"`
	ShmemGB             int `json:"shmem_gb"`
	SlabKB              int `json:"slab_kb"`
	SlabGB              int `json:"slab_gb"`
	SReclaimableKB      int `json:"s_reclaimable_kb"`
	SReclaimableGB      int `json:"s_reclaimable_gb"`
	SUnreclaimKB        int `json:"s_unreclaim_kb"`
	SUnreclaimGB        int `json:"s_unreclaim_gb"`
	KernelStackKB       int `json:"kernel_stack_kb"`
	KernelStackGB       int `json:"kernel_stack_gb"`
	PageTablesKB        int `json:"page_tables_kb"`
	PageTablesGB        int `json:"page_tables_gb"`
	NFSUnstableKB       int `json:"nfs_unstable_kb"`
	NFSUnstableGB       int `json:"nfs_unstable_gb"`
	BounceKB            int `json:"bounce_kb"`
	BounceGB            int `json:"bounce_gb"`
	WritebackTmpKB      int `json:"writeback_tmp_kb"`
	WritebackTmpGB      int `json:"writeback_tmp_gb"`
	CommitLimitKB       int `json:"commit_limit_kb"`
	CommitLimitGB       int `json:"commit_limit_gb"`
	CommittedASKB       int `json:"committed_a_s_kb"`
	CommittedASGB       int `json:"committed_a_s_gb"`
	VmallocTotalKB      int `json:"vmalloc_total_kb"`
	VmallocTotalGB      int `json:"vmalloc_total_gb"`
	VmallocUsedKB       int `json:"vmalloc_used_kb"`
	VmallocUsedGB       int `json:"vmalloc_used_gb"`
	VmallocChunkKB      int `json:"vmalloc_chunk_kb"`
	VmallocChunkGB      int `json:"vmalloc_chunk_gb"`
	HardwareCorruptedKB int `json:"hardware_corrupted_kb"`
	HardwareCorruptedGB int `json:"hardware_corrupted_gb"`
	AnonHugePagesKB     int `json:"anon_huge_pages_kb"`
	AnonHugePagesGB     int `json:"anon_huge_pages_gb"`
	HugePagesTot        int `json:"huge_pages_tot"`
	HugePagesFree       int `json:"huge_pages_free"`
	HugePagesRsvd       int `json:"huge_pages_rsvd"`
	HugePagesSurp       int `json:"huge_pages_surp"`
	HugePageSizeKB      int `json:"huge_pages_size_kb"`

	Last time.Time `json:"last"`
	TTL  int       `json:"ttl_sec"`
}

// Get memory info.
func (m *memory) get() error {
	o, err := common.LoadFileFields("/proc/meminfo", ":", []string{
		"MemTotal",
		"MemFree",
		"MemAvailable",
		"Cached",
		"SwapCached",
		"Active",
		"Inactive",
		"Unevictable",
		"Mlocked",
		"SwapTotal",
		"SwapFree",
		"Dirty",
		"Writeback",
		"AnonPages",
		"Mapped",
		"Shmem",
		"Slab",
		"SReclaimable",
		"SUnreclaim",
		"KernelStack",
		"PageTables",
		"NFS_Unstable",
		"Bounce",
		"WritebackTmp",
		"CommitLimit",
		"Committed_AS",
		"VmallocTotal",
		"VmallocUsed",
		"VmallocChunk",
		"HardwareCorrupted",
		"AnonHugePages",
		"HugePages_Total",
		"HugePages_Free",
		"HugePages_Rsvd",
		"HugePages_Surp",
		"Hugepagesize",
	})
	if err != nil {
		return err
	}

	// MemTotal
	m.TotalKB, err = strconv.Atoi(strings.TrimRight(o["MemTotal"], " kB"))
	if err != nil {
		return err
	}
	m.TotalGB = m.TotalKB / 1024 / 1024

	// MemFree
	m.FreeKB, err = strconv.Atoi(strings.TrimRight(o["MemFree"], " kB"))
	if err != nil {
		return err
	}
	m.FreeGB = m.FreeKB / 1024 / 1024

	// MemAvailable
	m.AvailableKB, err = strconv.Atoi(strings.TrimRight(o["MemAvailable"], " kB"))
	if err != nil {
		return err
	}
	m.AvailableGB = m.AvailableKB / 1024 / 1024

	// Cached
	m.CachedKB, err = strconv.Atoi(strings.TrimRight(o["Cached"], " kB"))
	if err != nil {
		return err
	}
	m.CachedGB = m.CachedKB / 1024 / 1024

	// SwapCached
	m.SwapCachedKB, err = strconv.Atoi(strings.TrimRight(o["SwapCached"], " kB"))
	if err != nil {
		return err
	}
	m.SwapCachedGB = m.SwapCachedKB / 1024 / 1024

	// Active
	m.ActiveKB, err = strconv.Atoi(strings.TrimRight(o["Active"], " kB"))
	if err != nil {
		return err
	}
	m.ActiveGB = m.ActiveKB / 1024 / 1024

	// Inactive
	m.InactiveKB, err = strconv.Atoi(strings.TrimRight(o["Inactive"], " kB"))
	if err != nil {
		return err
	}
	m.InactiveGB = m.InactiveKB / 1024 / 1024

	// Unevictable
	m.UnevictableKB, err = strconv.Atoi(strings.TrimRight(o["Unevictable"], " kB"))
	if err != nil {
		return err
	}
	m.UnevictableGB = m.UnevictableKB / 1024 / 1024

	// Mlocked
	m.MLockedKB, err = strconv.Atoi(strings.TrimRight(o["Mlocked"], " kB"))
	if err != nil {
		return err
	}
	m.MLockedGB = m.MLockedKB / 1024 / 1024

	// SwapTotal
	m.SwapTotalKB, err = strconv.Atoi(strings.TrimRight(o["SwapTotal"], " kB"))
	if err != nil {
		return err
	}
	m.SwapTotalGB = m.SwapTotalKB / 1024 / 1024

	// SwapFree
	m.SwapFreeKB, err = strconv.Atoi(strings.TrimRight(o["SwapFree"], " kB"))
	if err != nil {
		return err
	}
	m.SwapFreeGB = m.SwapFreeKB / 1024 / 1024

	// Dirty
	m.DirtyKB, err = strconv.Atoi(strings.TrimRight(o["Dirty"], " kB"))
	if err != nil {
		return err
	}
	m.DirtyGB = m.DirtyKB / 1024 / 1024

	// Writeback
	m.WritebackKB, err = strconv.Atoi(strings.TrimRight(o["Writeback"], " kB"))
	if err != nil {
		return err
	}
	m.WritebackGB = m.WritebackKB / 1024 / 1024

	// AnonPages
	m.AnonPagesKB, err = strconv.Atoi(strings.TrimRight(o["AnonPages"], " kB"))
	if err != nil {
		return err
	}
	m.AnonPagesGB = m.AnonPagesKB / 1024 / 1024

	// Mapped
	m.MappedKB, err = strconv.Atoi(strings.TrimRight(o["Mapped"], " kB"))
	if err != nil {
		return err
	}
	m.MappedGB = m.MappedKB / 1024 / 1024

	// Shmem
	m.ShmemKB, err = strconv.Atoi(strings.TrimRight(o["Shmen"], " kB"))
	if err != nil {
		return err
	}
	m.ShmemGB = m.ShmemKB / 1024 / 1024

	// Slab
	m.SlabKB, err = strconv.Atoi(strings.TrimRight(o["Slab"], " kB"))
	if err != nil {
		return err
	}
	m.SlabGB = m.SlabKB / 1024 / 1024

	// SReclaimable
	m.SReclaimableKB, err = strconv.Atoi(strings.TrimRight(o["SReclaimable"], " kB"))
	if err != nil {
		return err
	}
	m.SReclaimableGB = m.SReclaimableKB / 1024 / 1024

	// SUnreclaim
	m.SUnreclaimKB, err = strconv.Atoi(strings.TrimRight(o["SUnreclaim"], " kB"))
	if err != nil {
		return err
	}
	m.SUnreclaimGB = m.SUnreclaimKB / 1024 / 1024

	// KernelStack
	m.KernelStackKB, err = strconv.Atoi(strings.TrimRight(o["KernelStack"], " kB"))
	if err != nil {
		return err
	}
	m.KernelStackGB = m.KernelStackKB / 1024 / 1024

	// PageTables
	m.PageTablesKB, err = strconv.Atoi(strings.TrimRight(o["PageTables"], " kB"))
	if err != nil {
		return err
	}
	m.PageTablesGB = m.PageTablesKB / 1024 / 1024

	// NFS_Unstable
	m.NFSUnstableKB, err = strconv.Atoi(strings.TrimRight(o["NFS_Unstable"], " kB"))
	if err != nil {
		return err
	}
	m.NFSUnstableGB = m.NFSUnstableKB / 1024 / 1024

	// Bounce
	m.BounceKB, err = strconv.Atoi(strings.TrimRight(o["Bounce"], " kB"))
	if err != nil {
		return err
	}
	m.BounceGB = m.BounceKB / 1024 / 1024

	// WritebackTmp
	m.WritebackTmpKB, err = strconv.Atoi(strings.TrimRight(o["WritebackTmp"], " kB"))
	if err != nil {
		return err
	}
	m.WritebackTmpGB = m.WritebackTmpKB / 1024 / 1024

	// CommitLimit
	m.CommitLimitKB, err = strconv.Atoi(strings.TrimRight(o["CommitLimit"], " kB"))
	if err != nil {
		return err
	}
	m.CommitLimitGB = m.CommitLimitKB / 1024 / 1024

	// Committed_AS
	m.CommittedASKB, err = strconv.Atoi(strings.TrimRight(o["Committed_AS"], " kB"))
	if err != nil {
		return err
	}
	m.CommittedASGB = m.CommittedASKB / 1024 / 1024

	// VmallocTotal
	m.VmallocTotalKB, err = strconv.Atoi(strings.TrimRight(o["VmallocTotal"], " kB"))
	if err != nil {
		return err
	}
	m.VmallocTotalGB = m.VmallocTotalKB / 1024 / 1024

	// VmallocUsed
	m.VmallocUsedKB, err = strconv.Atoi(strings.TrimRight(o["VmallocUsed"], " kB"))
	if err != nil {
		return err
	}
	m.VmallocUsedGB = m.VmallocUsedKB / 1024 / 1024

	// VmallocChunk
	m.VmallocChunkKB, err = strconv.Atoi(strings.TrimRight(o["VmallocChunk"], " kB"))
	if err != nil {
		return err
	}
	m.VmallocChunkGB = m.VmallocChunkKB / 1024 / 1024

	// HardwareCorrupted
	m.HardwareCorruptedKB, err = strconv.Atoi(strings.TrimRight(o["HardwareCorrupted"], " kB"))
	if err != nil {
		return err
	}
	m.HardwareCorruptedGB = m.HardwareCorruptedKB / 1024 / 1024

	// AnonHugePages
	m.AnonHugePagesKB, err = strconv.Atoi(strings.TrimRight(o["AnonHugePages"], " kB"))
	if err != nil {
		return err
	}
	m.AnonHugePagesGB = m.AnonHugePagesKB / 1024 / 1024

	// HugePages_Total
	m.HugePagesTot, err = strconv.Atoi(o["HugePages_Total"])
	if err != nil {
		return err
	}

	// HugePages_Free
	m.HugePagesFree, err = strconv.Atoi(o["HugePages_Free"])
	if err != nil {
		return err
	}

	// HugePages_Rsvd
	m.HugePagesRsvd, err = strconv.Atoi(o["HugePages_Rsvd"])
	if err != nil {
		return err
	}

	// Hugepagesize
	m.HugePageSizeKB, err = strconv.Atoi(strings.TrimRight(o["Hugepagesize"], " kB"))
	if err != nil {
		return err
	}

	return nil
}
