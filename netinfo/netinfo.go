package netinfo

import (
	"github.com/mickep76/hwinfo/common"
	"net"
	"runtime"
)

// Info structure for information about a systems network interfaces.
type Interface struct {
	Name            string   `json:"name"`
	MTU             int      `json:"mtu"`
	IPAddr          []string `json:"ipaddr"`
	HWAddr          string   `json:"hwaddr"`
	Flags           []string `json:"flags"`
	Driver          string   `json:"driver,omitempty"`
	DriverVersion   string   `json:"driver_version,omitempty"`
	FirmwareVersion string   `json:"firmware_version,omitempty"`
}

// Info structure for information about a systems network.
type Info struct {
	Interfaces    []Interface `json:"interfaces"`
	OnloadVersion string      `json:"onload_version,omitempty"`
}

// GetInfo return information about a systems memory.
func GetInfo() (Info, error) {
	fields := []string{
		"driver",
		"version",
		"firmware-version",
	}

	info := Info{}

	intfs, err := net.Interfaces()
	if err != nil {
		return Info{}, err
	}

	for _, intf := range intfs {
		//		if intf.Name == "lo" || intf.Name == "lo0" {
		//			continue
		//		}

		addrs, err := intf.Addrs()
		if err != nil {
			return Info{}, err
		}

		nintf := Interface{
			Name:   intf.Name,
			HWAddr: intf.HardwareAddr.String(),
			MTU:    intf.MTU,
		}

		for _, addr := range addrs {
			nintf.IPAddr = append(nintf.IPAddr, addr.String())
		}

		switch runtime.GOOS {
		case "linux":
			o, err := common.ExecCmdFields("/usr/sbin/ethtool", []string{"-i", intf.Name}, ":", fields)
			if err != nil {
				return Info{}, err
			}

			nintf.Driver = o["driver"]
			nintf.DriverVersion = o["driver"]
			nintf.FirmwareVersion = o["firmware-version"]
		}

		info.Interfaces = append(info.Interfaces, nintf)
	}

	switch runtime.GOOS {
	case "linux":
		o, err := common.ExecCmdFields("/usr/bin/onload", []string{"--version"}, ":", []string{"Kernel module"})
		if err != nil {
			return Info{}, err
		}

		info.OnloadVersion = o["Kernel module"]
	}

	return info, nil
}
