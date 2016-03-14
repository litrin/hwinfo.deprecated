// +build darwin

package opsys

import (
	"runtime"

	"github.com/mickep76/hwinfo/common"
)

func (op *opSys) Get() error {
	o, err := common.ExecCmdFields("/usr/bin/sw_vers", []string{}, ":", []string{
		"ProductName",
		"ProductVersion",
	})
	if err != nil {
		return err
	}

	op.Kernel = runtime.GOOS
	op.Product = o["ProductName"]
	op.ProductVersion = o["ProductVersion"]

	op.KernelVersion, err = common.ExecCmd("uname", []string{"-r"})
	if err != nil {
		return err
	}

	return nil
}
