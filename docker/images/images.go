// +build linux

package images

import (
	//	"fmt"
	"os/exec"
	"regexp"
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
	ID       string `json:"Id"`
	Repo     string `json:"repo"`
	Tag      string `json:"tag"`
	Created  string `json:"created"`
	VirtSize string `json:"virtSize"`
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

func (cs *containers) ForceUpdate() error {
	cs.cache.LastUpdated = time.Now()
	cs.cache.FromCache = false

	o, err := exec.Command("docker", "images", "--no-trunc=true").Output()
	if err != nil {
		//		return err
		return nil
	}

	for c, line := range strings.Split(string(o), "\n") {
		re := regexp.MustCompile(`\s{2,}`)
		v := re.Split(line, -1)

		if c < 1 || len(v) < 5 {
			continue
		}

		d := DataItem{}
		d.Repo = v[0]
		d.Tag = v[1]
		d.ID = v[2]
		d.Created = v[3]
		d.VirtSize = v[4]

		*cs.data = append(*cs.data, d)
	}
	return nil
}
