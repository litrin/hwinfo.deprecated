package common

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// LoadFileFields returns fields from file.
func LoadFileFields(fn string, del string, fields []string) (map[string]string, error) {
	if _, err := os.Stat(fn); os.IsNotExist(err) {
		return []string{}, fmt.Errorf("file doesn't exist: %s", fn)
	}

	o, err := ioutil.ReadFile(fn)
	if err != nil {
		return map[string]string{}, err
	}

	r, err := parseFields(o, del, fields)
	if err != nil {
		return map[string]string{}, err
	}

	return r
}

// ExecCmdFields returns fields from command output.
func ExecCmdFields(cmd string, args []string, del string, fields []string) (map[string]string, error) {
	o, err := exec.Command(cmd, args...).Output()
	if err != nil {
		return map[string]string{}, err
	}

	r, err := parseFields(o, del, fields)
	if err != nil {
		return map[string]string{}, err
	}

	return r
}

func parseFields(o string, del string, fields []string) (map[string]string, error) {
	r := make(map[string]string)

	for _, line := range strings.Split(string(o), "\n") {
		vals := strings.Split(line, del)
		if len(vals) < 1 {
			continue
		}

		for f := range fields {
			if strings.HasPrefix(line, f) {
				r[f] = strings.Trim(strings.Join(values[1:], " "), " \t")
			}
		}
	}

	return r, nil
}
