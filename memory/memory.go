package memory

import (
	"time"
)

type Memory interface {
	Get() error
}

type Cached interface {
	SetTimeout(int)
	Get() error
	GetRefresh() error
}

type cached struct {
	Memory      *memory   `json:"memory"`
	LastUpdated time.Time `json:"last_updated"`
	Timeout     int       `json:"timeout_sec"`
	FromCache   bool      `json:"from_cache"`
}

func New() *memory {
	return &memory{}
}

func NewCached() *cached {
	return &cached{
		Memory:  New(),
		Timeout: 5 * 60, // 5 minutes
	}
}

func (c *cached) Get() error {
	if c.LastUpdated.IsZero() {
		if err := c.GetRefresh(); err != nil {
			return err
		}
	} else {
		expire := c.LastUpdated.Add(time.Duration(c.Timeout) * time.Second)
		if expire.Before(time.Now()) {
			if err := c.GetRefresh(); err != nil {
				return err
			}
		} else {
			c.FromCache = true
		}
	}

	return nil
}

func (c *cached) GetRefresh() error {
	if err := c.Memory.Get(); err != nil {
		return err
	}
	c.LastUpdated = time.Now()
	c.FromCache = false

	return nil
}
