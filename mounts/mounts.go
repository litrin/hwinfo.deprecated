// +build linux

package mounts

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

type Mounts interface {
	Get() error
}

type Cached interface {
	SetTimeout(int)
	Get() error
	GetRefresh() error
}

type mounts []mount

type mount struct {
	Source  string `json:"source"`
	Target  string `json:"target"`
	FSType  string `json:"fs_type"`
	Options string `json:"options"`
}

type cached struct {
	Mounts      *mounts   `json:"mounts"`
	LastUpdated time.Time `json:"last_updated"`
	Timeout     int       `json:"timeout_sec"`
	FromCache   bool      `json:"from_cache"`
}

func New() *mounts {
	return &mounts{}
}

func NewCached() *cached {
	return &cached{
		Mounts:  New(),
		Timeout: 5 * 60, // 5 minutes
	}
}

func (mounts *mounts) Get() error {
	fn := "/proc/mounts"
	if _, err := os.Stat(fn); os.IsNotExist(err) {
		return fmt.Errorf("file doesn't exist: %s", fn)
	}

	o, err := ioutil.ReadFile(fn)
	if err != nil {
		return err
	}

	for c, line := range strings.Split(string(o), "\n") {
		v := strings.Fields(line)
		if c < 1 || len(v) < 1 {
			continue
		}

		m := mount{}

		m.Source = v[0]
		m.Target = v[1]
		m.FSType = v[2]
		m.Options = v[3]

		*mounts = append(*mounts, m)
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
	if err := c.Mounts.Get(); err != nil {
		return err
	}
	c.LastUpdated = time.Now()
	c.FromCache = false

	return nil
}
