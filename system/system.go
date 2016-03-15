package system

import (
	"time"
)

type System interface {
	GetData() data
	GetCache() cache
	SetTimeout(int)
	Update() error
	ForceUpdate() error
}

type system struct {
	data  *data  `json:"data"`
	cache *cache `json:"cache"`
}

type cache struct {
	LastUpdated time.Time `json:"last_updated"`
	Timeout     int       `json:"timeout_sec"`
	FromCache   bool      `json:"from_cache"`
}

func New() System {
	return &system{
		data: &data{},
		cache: &cache{
			Timeout: 5 * 60, // 5 minutes
		},
	}
}

func (s *system) GetData() data {
	return *s.data
}

func (s *system) GetCache() cache {
	return *s.cache
}

func (s *system) SetTimeout(timeout int) {
	s.cache.Timeout = timeout
}

func (s *system) Update() error {
	if s.cache.LastUpdated.IsZero() {
		if err := s.ForceUpdate(); err != nil {
			return err
		}
	} else {
		expire := s.cache.LastUpdated.Add(time.Duration(s.cache.Timeout) * time.Second)
		if expire.Before(time.Now()) {
			if err := s.ForceUpdate(); err != nil {
				return err
			}
		} else {
			s.cache.FromCache = true
		}
	}

	return nil
}
