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

type Dock2box interface {
	SetTTL(int)
	Get() error
	Refresh() error
}

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
	FirstBoot string    `json:"firstboot"`
	CFlags    string    `json:"cflags_march_native"`
	Layers    []layer   `json:"layers"`
	Last      time.Time `json:"last"`
	TTL       int       `json:"ttl_sec"`
	Fresh     bool      `json:"fresh"`
}

// New constructor.
func New() *dock2box {
	return &dock2box{
		TTL: 12 * 60 * 60,
	}
}

// Get info.
func (d *dock2box) Get() error {
	if d.Last.IsZero() {
		if err := d.Refresh(); err != nil {
			return err
		}
	} else {
		expire := d.Last.Add(time.Duration(d.TTL) * time.Second)
		if expire.Before(time.Now()) {
			if err := d.Refresh(); err != nil {
				return err
			}
		} else {
			d.Fresh = false
		}
	}

	return nil
}

// Refresh cache.
func (d *dock2box) Refresh() error {
	if err := d.get(); err != nil {
		return err
	}
	d.Last = time.Now()
	d.Fresh = true

	return nil
}

func (d2b *dock2box) get() error {
	file := "/etc/dock2box/firstboot.json"
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return fmt.Errorf("file doesn't exist: %s", file)
	}

	o, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(o, &d2b); err != nil {
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

		d2b.Layers = append(d2b.Layers, l)
	}

	return nil
}
