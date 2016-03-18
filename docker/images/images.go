// +build linux

package images

import (
	"os/exec"
	"regexp"
	"strings"
	"time"
)

type Images interface {
	GetData() Data
	GetCache() Cache
	GetDataIntf() interface{}
	GetCacheIntf() interface{}
	SetTimeout(int)
	Update() error
	ForceUpdate() error
}

type images struct {
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

func New() Images {
	return &images{
		data: &Data{},
		cache: &Cache{
			Timeout: 5 * 60, // 5 minutes
		},
	}
}

func (i *images) GetData() Data {
	return *i.data
}

func (i *images) GetCache() Cache {
	return *i.cache
}

func (i *images) GetDataIntf() interface{} {
	return *i.data
}

func (i *images) GetCacheIntf() interface{} {
	return *i.cache
}

func (i *images) SetTimeout(timeout int) {
	i.cache.Timeout = timeout
}

func (i *images) Update() error {
	if i.cache.LastUpdated.IsZero() {
		if err := i.ForceUpdate(); err != nil {
			return err
		}
	} else {
		expire := i.cache.LastUpdated.Add(time.Duration(i.cache.Timeout) * time.Second)
		if expire.Before(time.Now()) {
			if err := i.ForceUpdate(); err != nil {
				return err
			}
		} else {
			i.cache.FromCache = true
		}
	}

	return nil
}

func (cs *images) ForceUpdate() error {
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
