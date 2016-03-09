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
		TTL: 5,
	}
}

// Get memory info.
func (m *memory) Get() {
	if m.Last.IsZero() {
		m.get()
		m.Last = time.Now()
	} else {
		expire := m.Last.Add(time.Duration(m.TTL) * time.Second)
		if expire.Before(time.Now()) {
			m.get()
		}
	}
}
