package memory

import (
	"time"
)

type Memory interface {
	SetTTL(int)
	Get() error
}

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
	TotalKB            int `json:"total_kb,omitempty"`
	TotalGB            int `json:"total_gb,omitempty"`
	FreeKB             int `json:"free_kb,omitempty"`
	FreeGB             int `json:"free_gb,omitempty"`
	AvailableKB        int `json:"available_kb,omitempty"`
	AvailableGB        int `json:"available_gb,omitempty"`
	CachedKB           int `json:"cached_kb,omitempty"`
	CachedGB           int `json:"cached_gb,omitempty"`
	SwapCachedKB       int `json:"swap_cached_kb,omitempty"`
	SwapCachedGB       int `json:"swap_cached_gb,omitempty"`
	CommittedActSizeKB int `json:"committed_act_size_kb,omitempty"`
	CommittedActSizeGB int `json:"committed_act_size_gb,omitempty"`
	HugePagesTot       int `json:"huge_pages_tot,omitempty"`
	HugePagesFree      int `json:"huge_pages_free,omitempty"`
	HugePagesRsvd      int `json:"huge_pages_rsvd,omitempty"`
	HugePagesSurp      int `json:"huge_pages_surp,omitempty"`
	HugePageSizeKB     int `json:"huge_pages_size_kb,omitempty"`

	last time.Time     `json:"-"`
	ttl  time.Duration `json:"-"`
}

// New memory constructor.
func New() *memory {
	return &memory{
		ttl: time.Duration(5) * time.Second,
	}
}

// Get memory info.
func (m *memory) Get() {
	if m.last.IsZero() {
		m.get()
		m.last = time.Now()
	} else {
		expire := m.last.Add(m.ttl)
		if expire.Before(time.Now()) {
			m.get()
		}
	}
}
