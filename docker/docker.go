// +build linux

package docker

import (
	"os/exec"
	"strings"
	"time"
)

type Docker interface {
	GetData() Data
	GetCache() Cache
	GetDataIntf() interface{}
	GetCacheIntf() interface{}
	SetTimeout(int)
	Update() error
	ForceUpdate() error
}

type docker struct {
	data  *Data  `json:"data"`
	cache *Cache `json:"cache"`
}

type Data struct {
	Version string `json:"version"`
	Build   string `json:"build"`
	Running bool   `json:"running"`
}

type Cache struct {
	LastUpdated time.Time `json:"last_updated"`
	Timeout     int       `json:"timeout_sec"`
	FromCache   bool      `json:"from_cache"`
}

func New() Docker {
	return &docker{
		data: &Data{},
		cache: &Cache{
			Timeout: 30, // 30 seconds
		},
	}
}

func (d *docker) GetData() Data {
	return *d.data
}

func (d *docker) GetCache() Cache {
	return *d.cache
}

func (d *docker) GetDataIntf() interface{} {
	return *d.data
}

func (d *docker) GetCacheIntf() interface{} {
	return *d.cache
}

func (d *docker) SetTimeout(timeout int) {
	d.cache.Timeout = timeout
}

func (d *docker) Update() error {
	if d.cache.LastUpdated.IsZero() {
		if err := d.ForceUpdate(); err != nil {
			return err
		}
	} else {
		expire := d.cache.LastUpdated.Add(time.Duration(d.cache.Timeout) * time.Second)
		if expire.Before(time.Now()) {
			if err := d.ForceUpdate(); err != nil {
				return err
			}
		} else {
			d.cache.FromCache = true
		}
	}

	return nil
}

func (d *docker) ForceUpdate() error {
	d.cache.LastUpdated = time.Now()
	d.cache.FromCache = false

	o, err := exec.Command("docker", "--version").Output()
	if err != nil {
		return err
	}

	v := strings.Fields(string(o))

	d.data.Version = strings.TrimRight(v[2], ",")
	d.data.Build = v[4]

	if err := exec.Command("docker", "ps"); err != nil {
		d.data.Running = false
	} else {
		d.data.Running = true
	}

	return nil
}
