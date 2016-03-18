// +build linux

package containers

import (
	"os/exec"
	"strings"
	"time"
)

type Containers interface {
	GetData() Data
	GetCache() Cache
	GetDataIntf() interface{}
	GetCacheIntf() interface{}
	SetTimeout(int)
	Update() error
	ForceUpdate() error
}

type containers struct {
	data  *Data  `json:"data"`
	cache *Cache `json:"cache"`
}

type Data []DataItem

type DataItem struct {
	ID      string `json:"id"`
	Image   string `json:"image"`
	Command string `json:"command"`
	Created string `json:"created"`
	Status  string `json:"status"`
	Ports   string `json:"ports"`
	Names   string `json:"names"`
}

type Cache struct {
	LastUpdated time.Time `json:"last_updated"`
	Timeout     int       `json:"timeout_sec"`
	FromCache   bool      `json:"from_cache"`
}

func New() Containers {
	return &containers{
		data: &Data{},
		cache: &Cache{
			Timeout: 5 * 60, // 5 minutes
		},
	}
}

func (c *containers) GetData() Data {
	return *c.data
}

func (c *containers) GetCache() Cache {
	return *c.cache
}

func (c *containers) GetDataIntf() interface{} {
	return *c.data
}

func (c *containers) GetCacheIntf() interface{} {
	return *c.cache
}

func (c *containers) SetTimeout(timeout int) {
	c.cache.Timeout = timeout
}

func (c *containers) Update() error {
	if c.cache.LastUpdated.IsZero() {
		if err := c.ForceUpdate(); err != nil {
			return err
		}
	} else {
		expire := c.cache.LastUpdated.Add(time.Duration(c.cache.Timeout) * time.Second)
		if expire.Before(time.Now()) {
			if err := c.ForceUpdate(); err != nil {
				return err
			}
		} else {
			c.cache.FromCache = true
		}
	}

	return nil
}

func (conts *containers) ForceUpdate() error {
	conts.cache.LastUpdated = time.Now()
	conts.cache.FromCache = false

	o, err := exec.Command("docker", "ps", "-a", "--no-trunc=true").Output()
	if err != nil {
		//		return err
		return nil
	}

	for c, line := range strings.Split(string(o), "\n") {
		v := strings.Fields(line)
		if c < 2 || len(v) < 7 {
			continue
		}

		cont := DataItem{}

		cont.ID = v[0]
		cont.Image = v[1]
		cont.Command = v[2]
		cont.Created = v[3]
		cont.Status = v[4]
		cont.Ports = v[5]
		cont.Names = v[6]

		*conts.data = append(*conts.data, cont)
	}

	return nil
}
