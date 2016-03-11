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

		v := variable{}

		v.Key = vals[0]
		v.Value = vals[2]

		s.Variables = append(s.Variables, v)
	}

	return nil
}
