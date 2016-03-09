package memory

import "time"

type Memory interface {
	SetTTL(int)
	Get() error
}

type memory struct {
	TotalGB            int `json:"total_gb,omitempty""`
	FreeGB             int `json:"free_gb,omitempty""`
	AvailableGB        int `json:"available_gb,omitempty""`
	CachedGB           int `json:"cached_gb,omitempty""`
	CommittedActSizeGB int `json:"committed_act_size_gb,omitempty""`
	HugePagesTot       int `json:"huge_pages_tot,omitempty""`
	HugePagesFree      int `json:"huge_pages_free,omitempty""`
	HugePagesRsvd      int `json:"huge_pages_rsvd,omitempty""`
	HugePagesSurp      int `json:"huge_pages_surp,omitempty""`
	HugePageSizeKB     int `json:"huge_pages_size_kb,omitempty""`

	last time.Time `json:"-"`
	ttl  int       `json:"-"`
}

// New memory constructor.
func New() *Memory {
	return &Memory{
		ttl: 60 * 5,
	}
}

func (m *MemoryS) Get() {
	expire := m.last
	expire.Add(time.Duration(m.ttl) * time.Second)
	if expire.Before(time.Now()) {
		m.GetNoCache()
	}
}
