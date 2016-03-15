package system

import (
	"time"
)

type System interface {
	SetTimeout(int)
	Get() error
	Refresh() error
	Data() *data
	Cache() *cache
}

type envelope struct {
	data  *data  `json:"data"`
	cache *cache `json:"cache"`
}

type cache struct {
	LastUpdated time.Time `json:"last_updated"`
	Timeout     int       `json:"timeout_sec"`
	FromCache   bool      `json:"from_cache"`
}

func New() System {
	return &envelope{
		data: &data{},
		cache: &cache{
			Timeout: 5 * 60, // 5 minutes
		},
	}
}

func (e *envelope) Get() error {
	if e.cache.LastUpdated.IsZero() {
		if err := e.Refresh(); err != nil {
			return err
		}
	} else {
		expire := e.cache.LastUpdated.Add(time.Duration(e.cache.Timeout) * time.Second)
		if expire.Before(time.Now()) {
			if err := e.Refresh(); err != nil {
				return err
			}
		} else {
			e.cache.FromCache = true
		}
	}

	return nil
}

func (e *envelope) SetTimeout(timeout int) {
	e.cache.Timeout = timeout
}

func (e *envelope) Data() *data {
	return e.data
}

func (e *envelope) Cache() *cache {
	return e.cache
}
