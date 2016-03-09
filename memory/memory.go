package memory

import (
	"time"
)

type Memory interface {
	SetTTL(int)
	Get() error
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
