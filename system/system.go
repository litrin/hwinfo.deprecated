package system

import (
	"time"
)

type System interface {
	SetTTL(int)
	Get() error
}

// New system constructor.
func New() *system {
	return &system{
		TTL: 12 * 60 * 60,
	}
}

// Get system info.
func (s *system) Get() error {
	if s.Last.IsZero() {
		if err := s.get(); err != nil {
			return err
		}
		s.Last = time.Now()
	} else {
		expire := s.Last.Add(time.Duration(s.TTL) * time.Second)
		if expire.Before(time.Now()) {
			if err := s.get(); err != nil {
				return err
			}
		}
	}

	return nil
}
