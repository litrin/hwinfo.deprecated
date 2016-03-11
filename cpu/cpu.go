package cpu

import (
	"time"
)

type CPU interface {
	SetTTL(int)
	Get() error
	Refresh() error
}

type cpu struct {
	Model          string    `json:"model"`
	Flags          string    `json:"flags"`
	Logical        int       `json:"logical"`
	Physical       int       `json:"physical"`
	Sockets        int       `json:"sockets"`
	CoresPerSocket int       `json:"cores_per_socket"`
	ThreadsPerCore int       `json:"threads_per_core"`
	Last           time.Time `json:"last"`
	TTL            int       `json:"ttl_sec"`
	Fresh          bool      `json:"fresh"`
}

// New constructor.
func New() *cpu {
	return &cpu{
		TTL: 12 * 60 * 60,
	}
}

// Get info.
func (c *cpu) Get() error {
	if c.Last.IsZero() {
		if err := c.Refresh(); err != nil {
			return err
		}
	} else {
		expire := c.Last.Add(time.Duration(c.TTL) * time.Second)
		if expire.Before(time.Now()) {
			if err := c.Refresh(); err != nil {
				return err
			}
		} else {
			c.Fresh = false
		}
	}

	return nil
}

// Refresh cache.
func (c *cpu) Refresh() error {
	if err := c.get(); err != nil {
		return err
	}
	c.Last = time.Now()
	c.Fresh = true

	return nil
}
