// +build linux

package disks

import (
	"path/filepath"
	"strconv"
	"time"

	"github.com/mickep76/hwinfo/common"
)

type Disks interface {
	Get() error
}

type Cached interface {
	SetTimeout(int)
	Get() error
	GetRefresh() error
}

type disks []disk

type disk struct {
	Device string `json:"device"`
	Name   string `json:"name"`
	SizeKB int    `json:"size_kb"`
	SizeGB int    `json:"size_gb"`
}

type cached struct {
	Disks       *disks    `json:"disks"`
	LastUpdated time.Time `json:"last_updated"`
	Timeout     int       `json:"timeout_sec"`
	FromCache   bool      `json:"from_cache"`
}

func New() Disks {
	return &disks{}
}

func NewCached() Cached {
	return &cached{
		Disks:   New(),
		Timeout: 5 * 60, // 5 minutes
	}
}

func (disks *disks) Get() error {
	files, err := filepath.Glob("/sys/class/block/*")
	if err != nil {
		return err
	}

	for _, path := range files {
		o, err := common.LoadFiles([]string{
			filepath.Join(path, "dev"),
			filepath.Join(path, "size"),
		})
		if err != nil {
			return err
		}

		d := disk{}

		d.Name = filepath.Base(path)
		d.Device = o["dev"]

		d.SizeKB, err = strconv.Atoi(o["size"])
		if err != nil {
			return err
		}
		d.SizeKB = d.SizeKB * 512 / 1024
		d.SizeGB = d.SizeKB / 1024 / 1024

		*disks = append(*disks, d)
	}

	return nil
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
	if err := c.Disks.Get(); err != nil {
		return err
	}
	c.LastUpdated = time.Now()
	c.FromCache = false

	return nil
}

func (c *cached) SetTimeout(timeout int) {
	c.Timeout = timeout
}
