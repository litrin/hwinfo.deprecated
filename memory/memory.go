package memory

import (
	"time"
)

type Memory interface {
	GetData() data
	GetCache() cache
	SetTimeout(int)
	Update() error
	ForceUpdate() error
}

type memory struct {
	data  *data  `json:"data"`
	cache *cache `json:"cache"`
}

type cache struct {
	LastUpdated time.Time `json:"last_updated"`
	Timeout     int       `json:"timeout_sec"`
	FromCache   bool      `json:"from_cache"`
}

func New() Memory {
	return &memory{
		data: &data{},
		cache: &cache{
			Timeout: 5 * 60, // 5 minutes
		},
	}
}

func (m *memory) GetData() data {
	return *m.data
}

func (m *memory) GetCache() cache {
	return *m.cache
}

func (m *memory) SetTimeout(timeout int) {
	m.cache.Timeout = timeout
}

func (m *memory) Update() error {
	if m.cache.LastUpdated.IsZero() {
		if err := m.ForceUpdate(); err != nil {
			return err
		}
	} else {
		expire := m.cache.LastUpdated.Add(time.Duration(m.cache.Timeout) * time.Second)
		if expire.Before(time.Now()) {
			if err := m.ForceUpdate(); err != nil {
				return err
			}
		} else {
			m.cache.FromCache = true
		}
	}

	return nil
}
