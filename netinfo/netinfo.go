package netinfo

import (
	"github.com/mickep76/hwinfo/common"
	"net"
	"runtime"
)

// Info structure for information about a systems memory.
type Info struct {
	Name            string   `json:"name"`
	MTU             int      `json:"mtu"`
	IPAddr          []string `json:"ipaddr"`
	HWAddr          string   `json:"hwaddr"`
	Driver          string   `json:"driver,omitempty"`
	DriverVersion   string   `json:"driver_version,omitempty"`
	FirmwareVersion string   `json:"firmware_version,omitempty"`
}

// GetInfo return information about a systems memory.
func GetInfo() ([]Info, error) {
	fields := []string{
		"driver",
		"version",
		"firmware-version",
	}

	i := []Info{}

	ifs, err := net.Interfaces()
	if err != nil {
		return []Info{}, err
	}

	for _, v := range ifs {
		if v.Name == "lo" || v.Name == "lo0" {
			continue
		}

		addrs, err := v.Addrs()
		if err != nil {
			return []Info{}, err
		}

		ia := []string{}
		for _, addr := range addrs {
			ia = append(ia, addr.String())
		}

		switch runtime.GOOS {
		case "linux":
			o, err := common.ExecCmdFields("/usr/sbin/ethtool", []string{"-i", v.Name}, ":", fields)
			if err != nil {
				return []Info{}, err
			}

			i = append(i, Info{
				Name:            v.Name,
				HWAddr:          v.HardwareAddr.String(),
				MTU:             v.MTU,
				IPAddr:          ia,
				Driver:          o["driver"],
				DriverVersion:   o["version"],
				FirmwareVersion: o["firmware-version"],
			})
		case "darwin":
			i = append(i, Info{
				Name:   v.Name,
				HWAddr: v.HardwareAddr.String(),
				MTU:    v.MTU,
				IPAddr: ia,
			})
		}
	}

	return i, nil
}
