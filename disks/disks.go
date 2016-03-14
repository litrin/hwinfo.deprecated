// +build linux

package disk

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

func New() *disks {
	return &disks{}
}

func NewCached() *cached {
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
