package opsys

import (
	"time"
)

type OpSys interface {
	SetTTL(int)
	Get() error
	Refresh() error
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

// New constructor.
func New() *opSys {
	return &opSys{
		TTL: 12 * 60 * 60,
	}
}

// Get info.
func (op *opSys) Get() error {
	if op.Last.IsZero() {
		if err := op.Refresh(); err != nil {
			return err
		}
	} else {
		expire := op.Last.Add(time.Duration(op.TTL) * time.Second)
		if expire.Before(time.Now()) {
			if err := op.Refresh(); err != nil {
				return err
			}
		} else {
			op.Fresh = false
		}
	}

	return nil
}

// Refresh cache.
func (op *opSys) Refresh() error {
	if err := op.get(); err != nil {
		return err
	}
	op.Last = time.Now()
	op.Fresh = true

	return nil
}
