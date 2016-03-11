// +build linux

package opsys

import (
	"github.com/mickep76/hwinfo/common"
	"runtime"
)

func (op *opSys) get() error {
	o, err := common.ExecCmdFields("lsb_release", []string{"-a"}, ":", []string{
		"Distributor ID",
		"Release",
	})
	if err != nil {
		return err
	}

	op.Kernel = runtime.GOOS
	op.Product = o["Distributor ID"]
	op.ProductVersion = o["Release"]

	op.KernelVersion, err = common.ExecCmd("uname", []string{"-r"})
	if err != nil {
		return err
	}

	return nil
}
