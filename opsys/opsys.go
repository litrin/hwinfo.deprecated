package opsys

import (
	"time"
)

type OpSys interface {
	Get() error
}

type Cached interface {
	SetTimeout(int)
	Get() error
	GetRefresh() error
}

type opSys struct {
	Kernel         string    `json:"kernel"`
	KernelVersion  string    `json:"kernel_version"`
	Product        string    `json:"product"`
	ProductVersion string    `json:"product_version"`
	Last           time.Time `json:"last"`
	TTL            int       `json:"ttl_sec"`
	Fresh          bool      `json:"fresh"`
}

type cached struct {
	OpSys       *opSys    `json:"op_sys"`
	LastUpdated time.Time `json:"last_updated"`
	Timeout     int       `json:"timeout_sec"`
	FromCache   bool      `json:"from_cache"`
}

func New() *opSys {
	return &opSys{}
}

func NewCached() *cached {
	return &cached{
		OpSys:   New(),
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
	if err := c.OpSys.Get(); err != nil {
		return err
	}
	c.LastUpdated = time.Now()
	c.FromCache = false

	return nil
}
