// +build linux

package sysctl

import (
	"os/exec"
	"strings"
)

func (s *sysctl) get() error {
	o, err := exec.Command("sysctl", "-a").Output()
	if err != nil {
		return err
	}

	for _, line := range strings.Split(string(o), "\n") {
		vals := strings.Fields(line)
		if len(vals) < 3 {
			continue
		}

		sys := Sysctl{}

		sys.Key = vals[0]
		sys.Value = vals[2]

		s = append(s, sys)
	}

	return nil
}
