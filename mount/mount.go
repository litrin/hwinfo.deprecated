// +build linux

package mount

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Mount interface {
    SetTTL(int)
    Get() error
    Refresh() error
}

type mount {
	FileSystems			[]fileSystem `json:"file_systems"`
    Last                time.Time `json:"last"`
    TTL                 int       `json:"ttl_sec"`
    Fresh               bool      `json:"fresh"`
}

type fileSystem struct {
	Source  string `json:"source"`
	Target  string `json:"target"`
	FSType  string `json:"fs_type"`
	Options string `json:"options"`
}

// New constructor.
func New() *mount{
	return &mount{
		TTL: 5,
	}
}

// Get info.
func (m *mount) Get() error {
	if m.Last.IsZero() {
		if err := m.Refresh(); err != nil {
			return err
		}
	} else {
		expire := m.Last.Add(time.Duration(m.TTL) * time.Second)
		if expire.Before(time.Now()) {
			if err := m.Refresh(); err != nil {
				return err
			}
		} else {
			m.Fresh = false
		}
	}

	return nil
}

// Refresh cache.
func (m *mount) Refresh() error {
	if err := m.get(); err != nil {
		return err
	}
	m.Last = time.Now()
	m.Fresh = true

	return nil
}

func (m *mount) get() error {
	fn := "/proc/mounts"
	if _, err := os.Stat(fn); os.IsNotExist(err) {
		return fmt.Errorf("file doesn't exist: %s", fn)
	}

	o, err := ioutil.ReadFile(fn)
	if err != nil {
		return err
	}

	for c, line := range strings.Split(string(o), "\n") {
		vals := strings.Fields(line)
		if c < 1 || len(vals) < 1 {
			continue
		}

		fs := fileSystem{}

		fs.Source = vals[0]
		fs.Target = vals[1]
		fs.FSType = vals[2]
		fs.Options = vals[3]

		m.FileSystems = append(m.FileSystems, fs)
	}

	return  nil
}
