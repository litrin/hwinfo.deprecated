// +build linux

package disk

import (
	"github.com/mickep76/hwinfo/common"
	"path/filepath"
	"strconv"
)

func (disk *disk) get() error {
	files, err := filepath.Glob("/sys/class/block/*")
	if err != nil {
		return err
	}

	for _, path := range files {
		o, err := common.LoadFiles([]string{
			filepath.Join(path, "dev"),
			filepath.Join(path, "size"),
		})
		if err != nil {
			return err
		}

		d := device{}

		d.Name = filepath.Base(path)
		d.Device = o["dev"]

		d.SizeGB, err = strconv.Atoi(o["size"])
		if err != nil {
			return err
		}
		d.SizeGB = d.SizeGB * 512 / 1024 / 1024 / 1024

		disk.Devices = append(disk.Devices, d)
	}

	return nil
}
