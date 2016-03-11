package opsys

import (
	"time"
)

type OpSys interface {
	SetTTL(int)
	Get() error
}

type opSys struct {
	Kernel         string    `json:"kernel"`
	KernelVersion  string    `json:"kernel_version"`
	Product        string    `json:"product"`
	ProductVersion string    `json:"product_version"`
	Last           time.Time `json:"last"`
	TTL            int       `json:"ttl_sec"`
}

// New OpSys constructor.
func New() *opSys {
	return &opSys{
		TTL: 12 * 60 * 60,
	}
}

// Get OpSys info.
func (o *opSys) Get() error {
	if o.Last.IsZero() {
		if err := o.get(); err != nil {
			return err
		}
		o.Last = time.Now()
	} else {
		expire := o.Last.Add(time.Duration(o.TTL) * time.Second)
		if expire.Before(time.Now()) {
			if err := o.get(); err != nil {
				return err
			}
		}
	}

	return nil
}
