// +build linux

package memory

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/mickep76/hwinfo/common"
)

type memory struct {
	TotalKB             int       `json:"total_kb"`
	TotalGB             int       `json:"total_gb"`
	FreeKB              int       `json:"free_kb"`
	FreeGB              int       `json:"free_gb"`
	AvailableKB         int       `json:"available_kb"`
	AvailableGB         int       `json:"available_gb"`
	BuffersKB           int       `json:"buffers_kb"`
	BuffersGB           int       `json:"buffers_gb"`
	CachedKB            int       `json:"cached_kb"`
	CachedGB            int       `json:"cached_gb"`
	SwapCachedKB        int       `json:"swap_cached_kb"`
	SwapCachedGB        int       `json:"swap_cached_gb"`
	ActiveKB            int       `json:"active_kb"`
	ActiveGB            int       `json:"active_gb"`
	InactiveKB          int       `json:"inactive_kb"`
	InactiveGB          int       `json:"inactive_gb"`
	ActiveAnonKB        int       `json:"active_anon_kb"`
	ActiveAnonGB        int       `json:"active_anon_kb"`
	InactiveAnonKB      int       `json:"inactive_anon_buffers_kb"`
	InactiveAnonGB      int       `json:"inactive_anon_buffers_kb"`
	ActiveFileKB        int       `json:"active_file_kb"`
	ActiveFileGB        int       `json:"active_file_gb"`
	InactiveFileKB      int       `json:"inactive_file_kb"`
	InactiveFileGB      int       `json:"inactive_file_gb"`
	UnevictableKB       int       `json:"unevictable_kb"`
	UnevictableGB       int       `json:"unevictable_gb"`
	MLockedKB           int       `json:"m_locked_kb"`
	MLockedGB           int       `json:"m_locked_gb"`
	SwapTotalKB         int       `json:"swap_total_kb"`
	SwapTotalGB         int       `json:"swap_total_gb"`
	SwapFreeKB          int       `json:"swap_free_kb"`
	SwapFreeGB          int       `json:"swap_free_gb"`
	DirtyKB             int       `json:"dirty_kb"`
	DirtyGB             int       `json:"dirty_gb"`
	WritebackKB         int       `json:"writeback_kb"`
	WritebackGB         int       `json:"writeback_gb"`
	AnonPagesKB         int       `json:"anon_pages_kb"`
	AnonPagesGB         int       `json:"anon_pages_gb"`
	MappedKB            int       `json:"mapped_kb"`
	MappedGB            int       `json:"mapped_gb"`
	ShmemKB             int       `json:"shmem_kb"`
	ShmemGB             int       `json:"shmem_gb"`
	SlabKB              int       `json:"slab_kb"`
	SlabGB              int       `json:"slab_gb"`
	SReclaimableKB      int       `json:"s_reclaimable_kb"`
	SReclaimableGB      int       `json:"s_reclaimable_gb"`
	SUnreclaimKB        int       `json:"s_unreclaim_kb"`
	SUnreclaimGB        int       `json:"s_unreclaim_gb"`
	KernelStackKB       int       `json:"kernel_stack_kb"`
	KernelStackGB       int       `json:"kernel_stack_gb"`
	PageTablesKB        int       `json:"page_tables_kb"`
	PageTablesGB        int       `json:"page_tables_gb"`
	NFSUnstableKB       int       `json:"nfs_unstable_kb"`
	NFSUnstableGB       int       `json:"nfs_unstable_gb"`
	BounceKB            int       `json:"bounce_kb"`
	BounceGB            int       `json:"bounce_gb"`
	WritebackTmpKB      int       `json:"writeback_tmp_kb"`
	WritebackTmpGB      int       `json:"writeback_tmp_gb"`
	CommitLimitKB       int       `json:"commit_limit_kb"`
	CommitLimitGB       int       `json:"commit_limit_gb"`
	CommittedASKB       int       `json:"committed_a_s_kb"`
	CommittedASGB       int       `json:"committed_a_s_gb"`
	VmallocTotalKB      int       `json:"vmalloc_total_kb"`
	VmallocTotalGB      int       `json:"vmalloc_total_gb"`
	VmallocUsedKB       int       `json:"vmalloc_used_kb"`
	VmallocUsedGB       int       `json:"vmalloc_used_gb"`
	VmallocChunkKB      int       `json:"vmalloc_chunk_kb"`
	VmallocChunkGB      int       `json:"vmalloc_chunk_gb"`
	HardwareCorruptedKB int       `json:"hardware_corrupted_kb"`
	HardwareCorruptedGB int       `json:"hardware_corrupted_gb"`
	AnonHugePagesKB     int       `json:"anon_huge_pages_kb"`
	AnonHugePagesGB     int       `json:"anon_huge_pages_gb"`
	HugePagesTot        int       `json:"huge_pages_tot"`
	HugePagesFree       int       `json:"huge_pages_free"`
	HugePagesRsvd       int       `json:"huge_pages_rsvd"`
	HugePagesSurp       int       `json:"huge_pages_surp"`
	HugePageSizeKB      int       `json:"huge_pages_size_kb"`
	DirectMap4kKB       int       `json:"direct_map_4k_kb"`
	DirectMap2MKB       int       `json:"direct_map_2m_kb"`
	DirectMap1GKB       int       `json:"direct_map_1g_kb"`
	Last                time.Time `json:"last"`
	TTL                 int       `json:"ttl_sec"`
}

// Get memory info.
func (m *memory) get() error {
	o, err := common.LoadFileFields("/proc/meminfo", ":", []string{
		"MemTotal",
		"MemFree",
		"MemAvailable",
		"Buffers",
		"Cached",
		"SwapCached",
		"Active",
		"Inactive",
		"Active(anon)",
		"Inactive(anon)",
		"Active(file)",
		"Inactive(file)",
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
		"DirectMap4k",
		"DirectMap2M",
		"DirectMap1G",
	})
	if err != nil {
		return err
	}

	// MemTotal
	m.TotalKB, err = strconv.Atoi(strings.TrimRight(o["MemTotal"], " kB"))
	if err != nil {
		return fmt.Errorf("failed parsing field: %s error: %s", "MemTotal", err.Error())
	}
	m.TotalGB = m.TotalKB / 1024 / 1024

	// MemFree
	m.FreeKB, err = strconv.Atoi(strings.TrimRight(o["MemFree"], " kB"))
	if err != nil {
		return fmt.Errorf("failed parsing field: %s error: %s", "MemFree", err.Error())
	}
	m.FreeGB = m.FreeKB / 1024 / 1024

	// MemAvailable
	m.AvailableKB, err = strconv.Atoi(strings.TrimRight(o["MemAvailable"], " kB"))
	if err != nil {
		return fmt.Errorf("failed parsing field: %s error: %s", "MemAvailable", err.Error())
	}
	m.AvailableGB = m.AvailableKB / 1024 / 1024

	// Buffers
	m.BuffersKB, err = strconv.Atoi(strings.TrimRight(o["Buffers"], " kB"))
	if err != nil {
		return fmt.Errorf("failed parsing field: %s error: %s", "Buffers", err.Error())
	}
	m.BuffersGB = m.BuffersKB / 1024 / 1024

	// Cached
	m.CachedKB, err = strconv.Atoi(strings.TrimRight(o["Cached"], " kB"))
	if err != nil {
		return fmt.Errorf("failed parsing field: %s error: %s", "Cached", err.Error())
	}
	m.CachedGB = m.CachedKB / 1024 / 1024

	// SwapCached
	m.SwapCachedKB, err = strconv.Atoi(strings.TrimRight(o["SwapCached"], " kB"))
	if err != nil {
		return fmt.Errorf("failed parsing field: %s error: %s", "SwapCached", err.Error())
	}
	m.SwapCachedGB = m.SwapCachedKB / 1024 / 1024

	// Active
	m.ActiveKB, err = strconv.Atoi(strings.TrimRight(o["Active"], " kB"))
	if err != nil {
		return fmt.Errorf("failed parsing field: %s error: %s", "Active", err.Error())
	}
	m.ActiveGB = m.ActiveKB / 1024 / 1024

	// Inactive
	m.InactiveKB, err = strconv.Atoi(strings.TrimRight(o["Inactive"], " kB"))
	if err != nil {
		return fmt.Errorf("failed parsing field: %s error: %s", "Inactive", err.Error())
	}
	m.InactiveGB = m.InactiveKB / 1024 / 1024

	// Active(anon)
	m.ActiveAnonKB, err = strconv.Atoi(strings.TrimRight(o["Active(anon)"], " kB"))
	if err != nil {
		return fmt.Errorf("failed parsing field: %s error: %s", "Active(anon)", err.Error())
	}
	m.ActiveAnonGB = m.ActiveAnonKB / 1024 / 1024

	// Inactive(anon)
	m.InactiveAnonKB, err = strconv.Atoi(strings.TrimRight(o["Inactive(anon)"], " kB"))
	if err != nil {
		return fmt.Errorf("failed parsing field: %s error: %s", "Inactive(anon)", err.Error())
	}
	m.InactiveAnonGB = m.InactiveAnonKB / 1024 / 1024

	// Active(file)
	m.ActiveFileKB, err = strconv.Atoi(strings.TrimRight(o["Active(file)"], " kB"))
	if err != nil {
		return fmt.Errorf("failed parsing field: %s error: %s", "Active(file)", err.Error())
	}
	m.ActiveFileGB = m.ActiveFileKB / 1024 / 1024

	// Inactive(file)
	m.InactiveFileKB, err = strconv.Atoi(strings.TrimRight(o["Inactive(file)"], " kB"))
	if err != nil {
		return fmt.Errorf("failed parsing field: %s error: %s", "Inactive(file)", err.Error())
	}
	m.InactiveFileGB = m.InactiveFileKB / 1024 / 1024

	// Unevictable
	m.UnevictableKB, err = strconv.Atoi(strings.TrimRight(o["Unevictable"], " kB"))
	if err != nil {
		return fmt.Errorf("failed parsing field: %s error: %s", "Unevictable", err.Error())
	}
	m.UnevictableGB = m.UnevictableKB / 1024 / 1024

	// Mlocked
	m.MLockedKB, err = strconv.Atoi(strings.TrimRight(o["Mlocked"], " kB"))
	if err != nil {
		return fmt.Errorf("failed parsing field: %s error: %s", "Mlocked", err.Error())
	}
	m.MLockedGB = m.MLockedKB / 1024 / 1024

	// SwapTotal
	m.SwapTotalKB, err = strconv.Atoi(strings.TrimRight(o["SwapTotal"], " kB"))
	if err != nil {
		return fmt.Errorf("failed parsing field: %s error: %s", "SwapTotal", err.Error())
	}
	m.SwapTotalGB = m.SwapTotalKB / 1024 / 1024

	// SwapFree
	m.SwapFreeKB, err = strconv.Atoi(strings.TrimRight(o["SwapFree"], " kB"))
	if err != nil {
		return fmt.Errorf("failed parsing field: %s error: %s", "SwapFree", err.Error())
	}
	m.SwapFreeGB = m.SwapFreeKB / 1024 / 1024

	// Dirty
	m.DirtyKB, err = strconv.Atoi(strings.TrimRight(o["Dirty"], " kB"))
	if err != nil {
		return fmt.Errorf("failed parsing field: %s error: %s", "Dirty", err.Error())
	}
	m.DirtyGB = m.DirtyKB / 1024 / 1024

	// Writeback
	m.WritebackKB, err = strconv.Atoi(strings.TrimRight(o["Writeback"], " kB"))
	if err != nil {
		return fmt.Errorf("failed parsing field: %s error: %s", "Writeback", err.Error())
	}
	m.WritebackGB = m.WritebackKB / 1024 / 1024

	// AnonPages
	m.AnonPagesKB, err = strconv.Atoi(strings.TrimRight(o["AnonPages"], " kB"))
	if err != nil {
		return fmt.Errorf("failed parsing field: %s error: %s", "AnonPages", err.Error())
	}
	m.AnonPagesGB = m.AnonPagesKB / 1024 / 1024

	// Mapped
	m.MappedKB, err = strconv.Atoi(strings.TrimRight(o["Mapped"], " kB"))
	if err != nil {
		return fmt.Errorf("failed parsing field: %s error: %s", "Mapped", err.Error())
	}
	m.MappedGB = m.MappedKB / 1024 / 1024

	// Shmem
	m.ShmemKB, err = strconv.Atoi(strings.TrimRight(o["Shmem"], " kB"))
	if err != nil {
		return fmt.Errorf("failed parsing field: %s error: %s", "Shmem", err.Error())
	}
	m.ShmemGB = m.ShmemKB / 1024 / 1024

	// Slab
	m.SlabKB, err = strconv.Atoi(strings.TrimRight(o["Slab"], " kB"))
	if err != nil {
		return fmt.Errorf("failed parsing field: %s error: %s", "Slab", err.Error())
	}
	m.SlabGB = m.SlabKB / 1024 / 1024

	// SReclaimable
	m.SReclaimableKB, err = strconv.Atoi(strings.TrimRight(o["SReclaimable"], " kB"))
	if err != nil {
		return fmt.Errorf("failed parsing field: %s error: %s", "SReclaimable", err.Error())
	}
	m.SReclaimableGB = m.SReclaimableKB / 1024 / 1024

	// SUnreclaim
	m.SUnreclaimKB, err = strconv.Atoi(strings.TrimRight(o["SUnreclaim"], " kB"))
	if err != nil {
		return fmt.Errorf("failed parsing field: %s error: %s", "SUnreclaim", err.Error())
	}
	m.SUnreclaimGB = m.SUnreclaimKB / 1024 / 1024

	// KernelStack
	m.KernelStackKB, err = strconv.Atoi(strings.TrimRight(o["KernelStack"], " kB"))
	if err != nil {
		return fmt.Errorf("failed parsing field: %s error: %s", "KernelStack", err.Error())
	}
	m.KernelStackGB = m.KernelStackKB / 1024 / 1024

	// PageTables
	m.PageTablesKB, err = strconv.Atoi(strings.TrimRight(o["PageTables"], " kB"))
	if err != nil {
		return fmt.Errorf("failed parsing field: %s error: %s", "PageTables", err.Error())
	}
	m.PageTablesGB = m.PageTablesKB / 1024 / 1024

	// NFS_Unstable
	m.NFSUnstableKB, err = strconv.Atoi(strings.TrimRight(o["NFS_Unstable"], " kB"))
	if err != nil {
		return fmt.Errorf("failed parsing field: %s error: %s", "NFS_Unstable", err.Error())
	}
	m.NFSUnstableGB = m.NFSUnstableKB / 1024 / 1024

	// Bounce
	m.BounceKB, err = strconv.Atoi(strings.TrimRight(o["Bounce"], " kB"))
	if err != nil {
		return fmt.Errorf("failed parsing field: %s error: %s", "Bounce", err.Error())
	}
	m.BounceGB = m.BounceKB / 1024 / 1024

	// WritebackTmp
	m.WritebackTmpKB, err = strconv.Atoi(strings.TrimRight(o["WritebackTmp"], " kB"))
	if err != nil {
		return fmt.Errorf("failed parsing field: %s error: %s", "WritebackTmp", err.Error())
	}
	m.WritebackTmpGB = m.WritebackTmpKB / 1024 / 1024

	// CommitLimit
	m.CommitLimitKB, err = strconv.Atoi(strings.TrimRight(o["CommitLimit"], " kB"))
	if err != nil {
		return fmt.Errorf("failed parsing field: %s error: %s", "CommitLimit", err.Error())
	}
	m.CommitLimitGB = m.CommitLimitKB / 1024 / 1024

	// Committed_AS
	m.CommittedASKB, err = strconv.Atoi(strings.TrimRight(o["Committed_AS"], " kB"))
	if err != nil {
		return fmt.Errorf("failed parsing field: %s error: %s", "Committed_AS", err.Error())
	}
	m.CommittedASGB = m.CommittedASKB / 1024 / 1024

	// VmallocTotal
	m.VmallocTotalKB, err = strconv.Atoi(strings.TrimRight(o["VmallocTotal"], " kB"))
	if err != nil {
		return fmt.Errorf("failed parsing field: %s error: %s", "VmallocTotal", err.Error())
	}
	m.VmallocTotalGB = m.VmallocTotalKB / 1024 / 1024

	// VmallocUsed
	m.VmallocUsedKB, err = strconv.Atoi(strings.TrimRight(o["VmallocUsed"], " kB"))
	if err != nil {
		return fmt.Errorf("failed parsing field: %s error: %s", "VmallocUsed", err.Error())
	}
	m.VmallocUsedGB = m.VmallocUsedKB / 1024 / 1024

	// VmallocChunk
	m.VmallocChunkKB, err = strconv.Atoi(strings.TrimRight(o["VmallocChunk"], " kB"))
	if err != nil {
		return fmt.Errorf("failed parsing field: %s error: %s", "VmallocChunk", err.Error())
	}
	m.VmallocChunkGB = m.VmallocChunkKB / 1024 / 1024

	// HardwareCorrupted
	m.HardwareCorruptedKB, err = strconv.Atoi(strings.TrimRight(o["HardwareCorrupted"], " kB"))
	if err != nil {
		return fmt.Errorf("failed parsing field: %s error: %s", "HardwareCorrupted", err.Error())
	}
	m.HardwareCorruptedGB = m.HardwareCorruptedKB / 1024 / 1024

	// AnonHugePages
	m.AnonHugePagesKB, err = strconv.Atoi(strings.TrimRight(o["AnonHugePages"], " kB"))
	if err != nil {
		return fmt.Errorf("failed parsing field: %s error: %s", "AnonHugePages", err.Error())
	}
	m.AnonHugePagesGB = m.AnonHugePagesKB / 1024 / 1024

	// HugePages_Total
	m.HugePagesTot, err = strconv.Atoi(o["HugePages_Total"])
	if err != nil {
		return fmt.Errorf("failed parsing field: %s error: %s", "HugePages_Total", err.Error())
	}

	// HugePages_Free
	m.HugePagesFree, err = strconv.Atoi(o["HugePages_Free"])
	if err != nil {
		return fmt.Errorf("failed parsing field: %s error: %s", "HugePages_Free", err.Error())
	}

	// HugePages_Rsvd
	m.HugePagesRsvd, err = strconv.Atoi(o["HugePages_Rsvd"])
	if err != nil {
		return fmt.Errorf("failed parsing field: %s error: %s", "HugePages_Rsvd", err.Error())
	}

	// Hugepagesize
	m.HugePageSizeKB, err = strconv.Atoi(strings.TrimRight(o["Hugepagesize"], " kB"))
	if err != nil {
		return fmt.Errorf("failed parsing field: %s error: %s", "Hugepagesize", err.Error())
	}

	// DirectMap4k
	m.DirectMap4kKB, err = strconv.Atoi(strings.TrimRight(o["DirectMap4k"], " kB"))
	if err != nil {
		return fmt.Errorf("failed parsing field: %s error: %s", "DirectMap4k", err.Error())
	}

	// DirectMap2M
	m.DirectMap2MKB, err = strconv.Atoi(strings.TrimRight(o["DirectMap2M"], " kB"))
	if err != nil {
		return fmt.Errorf("failed parsing field: %s error: %s", "DirectMap2M", err.Error())
	}

	// DirectMap1G
	m.DirectMap1GKB, err = strconv.Atoi(strings.TrimRight(o["DirectMap1G"], " kB"))
	if err != nil {
		return fmt.Errorf("failed parsing field: %s error: %s", "DirectMap1G", err.Error())
	}

	return nil
}
