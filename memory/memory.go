package memory

import (
	"time"
)

type Memory interface {
	SetTTL(int)
	Get() error
	Refresh() error
}

// New constructor.
func New() *memory {
	return &memory{
		TTL: 5,
	}
}

// Get info.
func (m *memory) Get() error {
	if m.Last.IsZero() {
		if err := m.Refresh(); err != nil {
			return err
		}
	} else {
		expire := m.Last.Add(time.Duration(m.TTL) * time.Second)
		if expire.Before(time.Now()) {
			if err := m.Refresh(); err != nil {
				return err
			}
		} else {
			m.Fresh = false
		}
	}

	return nil
}

// Refresh cache.
func (m *memory) Refresh() error {
	if err := m.get(); err != nil {
		return err
	}
	m.Last = time.Now()
	m.Fresh = true

	return nil
}
