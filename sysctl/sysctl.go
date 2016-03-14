/*
// +build linux
*/

package sysctl

import (
	"os/exec"
	"strings"
	"time"
)

type Sysctl interface {
	Get() error
}

type Cached interface {
	SetTimeout(int)
	Get() error
	GetRefresh() error
}

type sysctls []sysctl

type sysctl struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type cached struct {
	Sysctl      *sysctls  `json:"sysctl"`
	LastUpdated time.Time `json:"last_updated"`
	Timeout     int       `json:"timeout_sec"`
	FromCache   bool      `json:"from_cache"`
}

func (sysctls *sysctls) Get() error {
	o, err := exec.Command("sysctl", "-a").Output()
	if err != nil {
		return err
	}

	for _, line := range strings.Split(string(o), "\n") {
		vals := strings.Fields(line)
		if len(vals) < 3 {
			continue
		}

		s := sysctl{}

		s.Key = vals[0]
		s.Value = vals[2]

		*sysctls = append(*sysctls, s)
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
	if err := c.Sysctl.Get(); err != nil {
		return err
	}
	c.LastUpdated = time.Now()
	c.FromCache = false

	return nil
}
