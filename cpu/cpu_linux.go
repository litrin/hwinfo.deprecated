// +build linux

package cpu

import (
	"errors"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// GetInfo return information about a systems CPU(s).
func GetInfo() (Info, error) {
	c := Info{}

	if _, err := os.Stat("/proc/cpuinfo"); os.IsNotExist(err) {
		return Info{}, errors.New("file doesn't exist: /proc/cpuinfo")
	}

	o, err := ioutil.ReadFile("/proc/cpuinfo")
	if err != nil {
		return Info{}, err
	}

	cpuID := -1
	cpuIDs := make(map[int]bool)
	c.CoresPerSocket = 0
	c.Logical = 0
	for _, line := range strings.Split(string(o), "\n") {
		vals := strings.Split(line, ":")
		if len(vals) < 1 {
			continue
		}

		v := strings.Trim(strings.Join(vals[1:], " "), " ")
		if c.Model == "" && strings.HasPrefix(line, "model name") {
			c.Model = v
		} else if c.Flags == "" && strings.HasPrefix(line, "flags") {
			c.Flags = v
		} else if c.CoresPerSocket == 0 && strings.HasPrefix(line, "cpu cores") {
			c.CoresPerSocket, err = strconv.Atoi(v)
			if err != nil {
				return Info{}, err
			}
		} else if strings.HasPrefix(line, "processor") {
			c.Logical++
		} else if strings.HasPrefix(line, "physical id") {
			cpuID, err = strconv.Atoi(v)
			if err != nil {
				return Info{}, err
			}
			cpuIDs[cpuID] = true
		}
	}
	c.Sockets = int(len(cpuIDs))
	c.Physical = c.Sockets * c.CoresPerSocket
	c.ThreadsPerCore = c.Logical / c.Sockets / c.CoresPerSocket

	return c, nil
}
