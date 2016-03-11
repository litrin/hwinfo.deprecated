package cpu

import (
	"time"
)

type CPU interface {
	SetTTL(int)
	Get() error
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
}

// New cpu constructor.
func New() *cpu {
	return &cpu{
		TTL: 5,
	}
}

// Get cpu info.
func (c *cpu) Get() error {
	if c.Last.IsZero() {
		if err := c.get(); err != nil {
			return err
		}
		c.Last = time.Now()
	} else {
		expire := c.Last.Add(time.Duration(c.TTL) * time.Second)
		if expire.Before(time.Now()) {
			if err := c.get(); err != nil {
				return err
			}
		}
	}

	return nil
}
