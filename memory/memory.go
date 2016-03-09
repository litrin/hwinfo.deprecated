package memory

import (
	"fmt"
	"time"
)

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
