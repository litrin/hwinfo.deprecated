package sysctl

import (
	"time"
)

type Sysctl interface {
	SetTTL(int)
	Get() error
	Refresh() error
}

type sysctl struct {
	Variables []variable `json:"variables"`
	Last      time.Time  `json:"last"`
	TTL       int        `json:"ttl_sec"`
	Fresh     bool       `json:"fresh"`
}

type variable struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// New constructor.
func New() *sysctl {
	return &sysctl{
		TTL: 5 * 60 * 60,
	}
}

// Get info.
func (s *sysctl) Get() error {
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
func (s *sysctl) Refresh() error {
	if err := s.get(); err != nil {
		return err
	}
	s.Last = time.Now()
	s.Fresh = true

	return nil
}
