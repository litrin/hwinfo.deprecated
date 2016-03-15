// +build linux

package cpu

import (
	"errors"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

func (e *envelope) Refresh() error {
	e.cache.LastUpdated = time.Now()
	e.cache.FromCache = false

	if _, err := os.Stat("/proc/cpuinfo"); os.IsNotExist(err) {
		return errors.New("file doesn't exist: /proc/cpuinfo")
	}

	o, err := ioutil.ReadFile("/proc/cpuinfo")
	if err != nil {
		return err
	}

	cpuID := -1
	cpuIDs := make(map[int]bool)
	e.data.CoresPerSocket = 0
	e.data.Logical = 0
	for _, line := range strings.Split(string(o), "\n") {
		v := strings.Split(line, ":")
		if len(v) < 1 {
			continue
		}

		v := strings.Trim(strings.Join(v[1:], " "), " ")
		if e.data.Model == "" && strings.HasPrefix(line, "model name") {
			e.data.Model = v
		} else if e.data.flags == "" && strings.HasPrefix(line, "flags") {
			e.data.Flags = v
		} else if e.data.coresPerSocket == 0 && strings.HasPrefix(line, "cpu cores") {
			e.data.CoresPerSocket, err = strconv.Atoi(v)
			if err != nil {
				return err
			}
		} else if strings.HasPrefix(line, "processor") {
			e.data.Logical++
		} else if strings.HasPrefix(line, "physical id") {
			cpuID, err = strconv.Atoi(v)
			if err != nil {
				return err
			}
			cpuIDs[cpuID] = true
		}
	}
	e.data.Sockets = int(len(cpuIDs))
	e.data.Physical = c.Sockets * c.CoresPerSocket
	e.data.ThreadsPerCore = e.data.Logical / e.data.Sockets / e.data.CoresPerSocket

	return nil
}
