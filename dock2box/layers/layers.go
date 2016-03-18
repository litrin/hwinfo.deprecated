// +build linux

package layers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

type Layers interface {
	GetData() Data
	GetCache() Cache
	GetDataIntf() interface{}
	GetCacheIntf() interface{}
	SetTimeout(int)
	Update() error
	ForceUpdate() error
}

type layers struct {
	data  *Data  `json:"data"`
	cache *Cache `json:"cache"`
}

type Data []DataItem

type DataItem struct {
	Layer             string `json:"layer"`
	Image             string `json:"image"`
	Repo              string `json:"repo"`
	Commit            string `json:"commit"`
	Created           string `json:"created"`
	CPU               string `json:"cpu"`
	CPUFlags          string `json:"cpuflags"`
	KernelConfig      string `json:"kernelconfig"`
	GCCVersion        string `json:"gcc_version"`
	CFlags            string `json:"cflags"`
	CFlagsMarchNative string `json:"cflags_march_native"`
}

type Cache struct {
	LastUpdated time.Time `json:"last_updated"`
	Timeout     int       `json:"timeout_sec"`
	FromCache   bool      `json:"from_cache"`
}

func New() Layers {
	return &layers{
		data: &Data{},
		cache: &Cache{
			Timeout: 5 * 60, // 5 minutes
		},
	}
}

func (l *layers) GetData() Data {
	return *l.data
}

func (l *layers) GetCache() Cache {
	return *l.cache
}

func (l *layers) GetDataIntf() interface{} {
	return *l.data
}

func (l *layers) GetCacheIntf() interface{} {
	return *l.cache
}

func (l *layers) SetTimeout(timeout int) {
	l.cache.Timeout = timeout
}

func (l *layers) Update() error {
	if l.cache.LastUpdated.IsZero() {
		if err := l.ForceUpdate(); err != nil {
			return err
		}
	} else {
		expire := l.cache.LastUpdated.Add(time.Duration(l.cache.Timeout) * time.Second)
		if expire.Before(time.Now()) {
			if err := l.ForceUpdate(); err != nil {
				return err
			}
		} else {
			l.cache.FromCache = true
		}
	}

	return nil
}

func (ls *layers) ForceUpdate() error {
	ls.cache.LastUpdated = time.Now()
	ls.cache.FromCache = false

	files, err := filepath.Glob("/etc/dock2box/layers/*.json")
	if err != nil {
		return err
	}

	for _, file := range files {
		if _, err := os.Stat(file); os.IsNotExist(err) {
			return fmt.Errorf("file doesn't exist: %s", file)
		}

		o, err := ioutil.ReadFile(file)
		if err != nil {
			return err
		}

		l := DataItem{}
		if err := json.Unmarshal(o, &l); err != nil {
			return err
		}

		fn := path.Base(file)
		l.Layer = strings.TrimSuffix(fn, filepath.Ext(fn))

		*ls.data = append(*ls.data, l)
	}

	return nil
}
