// +build linux

package lvm

import (
	"time"
)

type LVM interface {
	SetTTL(int)
	Get() error
	Refresh() error
}

type lvm struct {
	PhysVols []physVol `json:"phys_vols"`
	LogVols  []logVol  `json:"log_vols"`
	VolGrps  []volGrp  `json:"vol_grps"`
	Last     time.Time `json:"last"`
	TTL      int       `json:"ttl_sec"`
	Fresh    bool      `json:"fresh"`
}

type physVol struct {
	Name   string `json:"name"`
	VolGrp string `json:"vol_group"`
	Format string `json:"format"`
	Attr   string `json:"attr"`
	SizeGB int    `json:"size_gb"`
	FreeGB int    `json:"free_gb"`
}

type logVol struct {
	Name   string `json:"name"`
	VolGrp string `json:"vol_grp"`
	Attr   string `json:"attr"`
	SizeGB int    `json:"size_gb"`
}

type volGrp struct {
	Name   string `json:"name"`
	Attr   string `json:"attr"`
	SizeGB int    `json:"size_gb"`
	FreeGB int    `json:"free_gb"`
}

// New constructor.
func New() *lvm {
	return &lvm{
		TTL: 60 * 60,
	}
}

// Get info.
func (l *lvm) Get() error {
	if l.Last.IsZero() {
		if err := l.Refresh(); err != nil {
			return err
		}
	} else {
		expire := l.Last.Add(time.Duration(l.TTL) * time.Second)
		if expire.Before(time.Now()) {
			if err := l.Refresh(); err != nil {
				return err
			}
		} else {
			l.Fresh = false
		}
	}

	return nil
}

// Refresh cache.
func (l *lvm) Refresh() error {
	if err := l.get(); err != nil {
		return err
	}
	l.Last = time.Now()
	l.Fresh = true

	return nil
}
