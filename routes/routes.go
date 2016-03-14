package routes

import (
	"os/exec"
	"strconv"
	"strings"
	"time"
)

type Routes interface {
	Get() error
}

type Cached interface {
	SetTimeout(int)
	Get() error
	Refresh() error
}

type routes []route

type route struct {
	Destination string `json:"destination"`
	Gateway     string `json:"gateway"`
	Genmask     string `json:"genmask"`
	Flags       string `json:"flags"`
	MSS         int    `json:"mss"` // Maximum segment size
	Window      int    `json:"window"`
	IRTT        int    `json:"irtt"` // Initial round trip time
	Interface   string `json:"interface"`
}

type cached struct {
	Routes      *routes   `json:"routes"`
	LastUpdated time.Time `json:"last_updated"`
	Timeout     int       `json:"timeout_sec"`
	FromCache   bool      `json:"from_cache"`
}

func New() *routes {
	r := routes{}
	return &r
}

func NewCached() *cached {
	c := cached{
		Routes:  New(),
		Timeout: 5 * 60 * 60,
	}
	return &c
}

func (routes *routes) Get() error {
	o, err := exec.Command("netstat", "-rn").Output()
	if err != nil {
		return err
	}

	for c, line := range strings.Split(string(o), "\n") {
		v := strings.Fields(line)
		if c < 2 || len(v) < 8 {
			continue
		}

		r := route{}

		r.Destination = v[0]
		r.Gateway = v[1]
		r.Genmask = v[2]
		r.Flags = v[3]

		r.MSS, err = strconv.Atoi(v[4])
		if err != nil {
			return err
		}

		r.Window, err = strconv.Atoi(v[5])
		if err != nil {
			return err
		}

		r.IRTT, err = strconv.Atoi(v[6])
		if err != nil {
			return err
		}

		r.Interface = v[7]

		*routes = append(*routes, r)
	}

	return nil
}

func (c *cached) Get() error {
	if c.LastUpdated.IsZero() {
		if err := c.Refresh(); err != nil {
			return err
		}
	} else {
		expire := c.LastUpdated.Add(time.Duration(c.Timeout) * time.Second)
		if expire.Before(time.Now()) {
			if err := c.Refresh(); err != nil {
				return err
			}
		} else {
			c.FromCache = true
		}
	}

	return nil
}

func (c *cached) Refresh() error {
	if err := c.Routes.Get(); err != nil {
		return err
	}
	c.LastUpdated = time.Now()
	c.FromCache = false

	return nil
}
