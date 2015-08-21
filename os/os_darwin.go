// +build darwin

package os

import (
	"github.com/mickep76/hwinfo/common"
	"strconv"
)

// GetInfo return information about a systems memory.
func GetInfo() (Info, error) {
	fields := []string{
		"ProductName",
		"ProductVersion",
	}

	i := Info{}

	o, err := common.ExecCmdFields("/usr/sbin/sw_vers", []string{}, ":", fields)
	if err != nil {
		return Info{}, err
	}

	i.Product = o["ProductName"]
	i.Version = o["ProductVersion"]

	return r, nil
}
