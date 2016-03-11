package system

import (
	"time"
)

type System interface {
	SetTTL(int)
	Get() error
	Refresh() error
}

// New system constructor.
func New() *system {
	return &system{
		TTL: 12 * 60 * 60,
	}
}

// Get info.
func (s *system) Get() error {
	if s.Last.IsZero() {
		if err := s.Refresh(); err != nil {
			return err
		}
	} else {
		expire := s.Last.Add(time.Duration(s.TTL) * time.Second)
		if expire.Before(time.Now()) {
			if err := s.Refresh(); err != nil {
				return err
			}
		} else {
			s.Fresh = false
		}
	}

	return nil
}

// Refresh cache.
func (s *system) Refresh() error {
	if err := s.get(); err != nil {
		return err
	}
	s.Last = time.Now()
	s.Fresh = true

	return nil
}
