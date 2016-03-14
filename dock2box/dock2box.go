// +build linux

package dock2box

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

type Dock2Box interface {
	Get() error
}

type Cached interface {
	SetTimeout(int)
	Get() error
	GetRefresh() error
}

type layers []layer

type layer struct {
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

type dock2box struct {
	FirstBoot string  `json:"firstboot"`
	CFlags    string  `json:"cflags_march_native"`
	Layers    *layers `json:"layers"`
}

type cached struct {
	Dock2Box    *dock2box `json:"dock2box"`
	LastUpdated time.Time `json:"last_updated"`
	Timeout     int       `json:"timeout_sec"`
	FromCache   bool      `json:"from_cache"`
}

func New() *dock2box {
	return &dock2box{}
}

func NewCached() *cached {
	return &cached{
		Dock2Box: New(),
		Timeout:  12 * 60 * 60, // 12 hours
	}
}

func (d *dock2box) Get() error {
	file := "/etc/dock2box/firstboot.json"
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return fmt.Errorf("file doesn't exist: %s", file)
	}

	o, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(o, &d); err != nil {
		return err
	}

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

		l := layer{}
		if err := json.Unmarshal(o, &l); err != nil {
			return err
		}

		fn := path.Base(file)
		l.Layer = strings.TrimSuffix(fn, filepath.Ext(fn))

		*d.Layers = append(*d.Layers, l)
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
	if err := c.Dock2Box.Get(); err != nil {
		return err
	}
	c.LastUpdated = time.Now()
	c.FromCache = false

	return nil
}
