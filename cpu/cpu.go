package cpu

import (
	"time"
)

type CPU interface {
	Get() error
}

type Cached interface {
	SetTimeout(int)
	Get() error
	GetRefresh() error
}

type cpu struct {
	Model          string `json:"model"`
	Flags          string `json:"flags"`
	Logical        int    `json:"logical"`
	Physical       int    `json:"physical"`
	Sockets        int    `json:"sockets"`
	CoresPerSocket int    `json:"cores_per_socket"`
	ThreadsPerCore int    `json:"threads_per_core"`
}

type cached struct {
	CPU         *cpu      `json:"cpu"`
	LastUpdated time.Time `json:"last_updated"`
	Timeout     int       `json:"timeout_sec"`
	FromCache   bool      `json:"from_cache"`
}

func New() *cpu {
	return &cpu{}
}

func NewCached() *cached {
	return &cached{
		CPU:     New(),
		Timeout: 12 * 60 * 60, // 12 hours
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
	if err := c.CPU.Get(); err != nil {
		return err
	}
	c.LastUpdated = time.Now()
	c.FromCache = false

	return nil
}
