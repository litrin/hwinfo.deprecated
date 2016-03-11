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
func (m *memory) Get() error {
	if m.Last.IsZero() {
		if err := m.get(); err != nil {
			return err
		}
		m.Last = time.Now()
	} else {
		expire := m.Last.Add(time.Duration(m.TTL) * time.Second)
		if expire.Before(time.Now()) {
			if err := m.get(); err != nil {
				return err
			}
		}
	}

	return nil
}
