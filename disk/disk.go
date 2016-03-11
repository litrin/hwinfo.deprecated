package disk

import (
	"time"
)

type Disk interface {
	SetTTL(int)
	Get() error
	Refresh() error
}

type disk struct {
	devices []device  `json:"devices"`
	Last    time.Time `json:"last"`
	TTL     int       `json:"ttl_sec"`
	Fresh   bool      `json:"fresh"`
}

type device struct {
	Device string `json:"device"`
	Name   string `json:"name"`
	SizeGB int    `json:"size_gb"`
}

// New constructor.
func New() *disk {
	return &disk{
		TTL: 5 * 60,
	}
}

// Get info.
func (d *disk) Get() error {
	if d.Last.IsZero() {
		if err := d.Refresh(); err != nil {
			return err
		}
	} else {
		expire := d.Last.Add(time.Duration(d.TTL) * time.Second)
		if expire.Before(time.Now()) {
			if err := d.Refresh(); err != nil {
				return err
			}
		} else {
			d.Fresh = false
		}
	}

	return nil
}

// Refresh cache.
func (d *disk) Refresh() error {
	if err := d.get(); err != nil {
		return err
	}
	d.Last = time.Now()
	d.Fresh = true

	return nil
}
