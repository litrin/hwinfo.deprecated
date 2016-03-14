// +build linux

package interfaces

import (
	"fmt"
	"net"
	//	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/mickep76/hwinfo/common"
)

type Interfaces interface {
	Get() error
}

type Cached interface {
	SetTimeout(int)
	Get() error
	GetRefresh() error
}

type intfs []intf

type intf struct {
	Name            string   `json:"name"`
	MTU             int      `json:"mtu"`
	IPAddr          []string `json:"ipaddr"`
	HWAddr          string   `json:"hwaddr"`
	Flags           []string `json:"flags"`
	Driver          string   `json:"driver,omitempty"`
	DriverVersion   string   `json:"driver_version,omitempty"`
	FirmwareVersion string   `json:"firmware_version,omitempty"`
	PCIBus          string   `json:"pci_bus,omitempty"`
	PCIBusURL       string   `json:"pci_bus_url,omitempty"`
	SwChassisID     string   `json:"sw_chassis_id"`
	SwName          string   `json:"sw_name"`
	SwDescr         string   `json:"sw_descr"`
	SwPortID        string   `json:"sw_port_id"`
	SwPortDescr     string   `json:"sw_port_descr"`
	SwVLAN          string   `json:"sw_vlan"`
}

type cached struct {
	Interfaces  *intfs    `json:"interfaces"`
	LastUpdated time.Time `json:"last_updated"`
	Timeout     int       `json:"timeout_sec"`
	FromCache   bool      `json:"from_cache"`
}

func New() *intfs {
	return &intfs{}
}

func NewCached() *cached {
	return &cached{
		Interfaces: New(),
		Timeout:    5 * 60, // 5 minutes
	}
}

func (sIntfs *intfs) Get() error {
	rIntfs, err := net.Interfaces()
	if err != nil {
		return err
	}

	for _, rIntf := range rIntfs {
		// Skip loopback interfaces
		if rIntf.Flags&net.FlagLoopback != 0 {
			continue
		}

		addrs, err := rIntf.Addrs()
		if err != nil {
			return err
		}

		sIntf := intf{
			Name:   rIntf.Name,
			HWAddr: rIntf.HardwareAddr.String(),
			MTU:    rIntf.MTU,
		}

		for _, addr := range addrs {
			sIntf.IPAddr = append(sIntf.IPAddr, addr.String())
		}

		if rIntf.Flags&net.FlagUp != 0 {
			sIntf.Flags = append(sIntf.Flags, "up")
		}
		if rIntf.Flags&net.FlagBroadcast != 0 {
			sIntf.Flags = append(sIntf.Flags, "broadcast")
		}
		if rIntf.Flags&net.FlagPointToPoint != 0 {
			sIntf.Flags = append(sIntf.Flags, "pointtopoint")
		}
		if rIntf.Flags&net.FlagMulticast != 0 {
			sIntf.Flags = append(sIntf.Flags, "multicast")
		}

		switch runtime.GOOS {
		case "linux":
			o, err := common.ExecCmdFields("/usr/sbin/ethtool", []string{"-i", rIntf.Name}, ":", []string{
				"driver",
				"version",
				"firmware-version",
				"bus-info",
			})
			if err != nil {
				return err
			}

			sIntf.Driver = o["driver"]
			sIntf.DriverVersion = o["version"]
			sIntf.FirmwareVersion = o["firmware-version"]
			if strings.HasPrefix(o["bus-info"], "0000:") {
				sIntf.PCIBus = o["bus-info"]
				sIntf.PCIBusURL = fmt.Sprintf("/pci/%v", o["bus-info"])
			}

			o2, err := common.ExecCmdFields("/usr/sbin/lldpctl", []string{rIntf.Name}, ":", []string{
				"ChassisID",
				"SysName",
				"SysDescr",
				"PortID",
				"PortDescr",
				"VLAN",
			})
			if err != nil {
				return err
			}

			sIntf.SwChassisID = o2["ChassisID"]
			sIntf.SwName = o2["SysName"]
			sIntf.SwDescr = o2["SysDescr"]
			sIntf.SwPortID = o2["PortID"]
			sIntf.SwPortDescr = o2["PortDescr"]
			sIntf.SwVLAN = o2["VLAN"]
		}

		*sIntfs = append(*sIntfs, sIntf)
	}

	/*
		switch runtime.GOOS {
		case "linux":
			_, err := exec.LookPath("onload")
			if err == nil {
				o, _ := common.ExecCmdFields("onload", []string{"--version"}, ":", []string{"Kernel module"})
				if err != nil {
					return err
				}
				n.OnloadVersion = o["Kernel module"]
			}
		}
	*/

	return nil
}

func (c *cached) Get() error {
	if c.LastUpdated.IsZero() {
		if err := c.GetRefresh(); err != nil {
			return err
		}
	} else {
		expire := c.LastUpdated.Add(time.Duration(c.Timeout) * time.Second)
		if expire.Before(time.Now()) {
			if err := c.GetRefresh(); err != nil {
				return err
			}
		} else {
			c.FromCache = true
		}
	}

	return nil
}

func (c *cached) GetRefresh() error {
	if err := c.Interfaces.Get(); err != nil {
		return err
	}
	c.LastUpdated = time.Now()
	c.FromCache = false

	return nil
}
