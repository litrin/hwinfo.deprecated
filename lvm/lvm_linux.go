// +build linux

package lvm

import (
	"errors"
	"os/exec"
	"strconv"
	"strings"
)

func (l *lvm) get() error {
	pvs, err := GetPhysVols()
	if err != nil {
		return err
	}
	l.PhysVols = pvs

	lvs, err := GetLogVols()
	if err != nil {
		return err
	}
	l.LogVols = lvs

	vgs, err := GetVolGrps()
	if err != nil {
		return err
	}
	l.VolGrps = vgs

	return nil
}

func GetPhysVols() ([]physVol, error) {
	pvs := []physVol{}

	_, err := exec.LookPath("pvs")
	if err != nil {
		return []physVol{}, errors.New("command doesn't exist: pvs")
	}

	o, err := exec.Command("pvs", "--units", "B").Output()
	if err != nil {
		return []physVol{}, err
	}

	for c, line := range strings.Split(string(o), "\n") {
		vals := strings.Fields(line)
		if c < 1 || len(vals) < 1 {
			continue
		}

		pv := physVol{}

		pv.Name = vals[0]
		pv.VolGrp = vals[1]
		pv.Format = vals[2]
		pv.Attr = vals[3]

		pv.SizeGB, err = strconv.Atoi(strings.TrimRight(vals[4], "B"))
		if err != nil {
			return []physVol{}, err
		}
		pv.SizeGB = pv.SizeGB / 1024 / 1024 / 1024

		pv.FreeGB, err = strconv.Atoi(strings.TrimRight(vals[5], "B"))
		if err != nil {
			return []physVol{}, err
		}
		pv.FreeGB = pv.FreeGB / 1024 / 1024 / 1024

		pvs = append(pvs, pv)
	}

	return pvs, nil
}

func GetLogVols() ([]logVol, error) {
	lvs := []logVol{}

	_, err := exec.LookPath("lvs")
	if err != nil {
		return []logVol{}, errors.New("command doesn't exist: lvs")
	}

	o, err := exec.Command("lvs", "--units", "B").Output()
	if err != nil {
		return []logVol{}, err
	}

	for c, line := range strings.Split(string(o), "\n") {
		vals := strings.Fields(line)
		if c < 1 || len(vals) < 1 {
			continue
		}

		lv := logVol{}

		lv.Name = vals[0]
		lv.VolGrp = vals[1]
		lv.Attr = vals[2]

		lv.SizeGB, err = strconv.Atoi(strings.TrimRight(vals[3], "B"))
		if err != nil {
			return []logVol{}, err
		}
		lv.SizeGB = lv.SizeGB / 1024 / 1024 / 1024

		lvs = append(lvs, lv)
	}

	return lvs, nil
}

func GetVolGrps() ([]volGrp, error) {
	vgs := []volGrp{}

	_, err := exec.LookPath("vgs")
	if err != nil {
		return []volGrp{}, errors.New("command doesn't exist: vgs")
	}

	o, err := exec.Command("vgs", "--units", "B").Output()
	if err != nil {
		return []volGrp{}, err
	}

	for c, line := range strings.Split(string(o), "\n") {
		vals := strings.Fields(line)
		if c < 1 || len(vals) < 1 {
			continue
		}

		vg := volGrp{}

		vg.Name = vals[0]
		vg.Attr = vals[4]

		vg.SizeGB, err = strconv.Atoi(strings.TrimRight(vals[5], "B"))
		if err != nil {
			return []volGrp{}, err
		}
		vg.SizeGB = vg.SizeGB / 1024 / 1024 / 1024

		vg.FreeGB, err = strconv.Atoi(strings.TrimRight(vals[6], "B"))
		if err != nil {
			return []volGrp{}, err
		}
		vg.FreeGB = vg.FreeGB / 1024 / 1024 / 1024

		vgs = append(vgs, vg)
	}

	return vgs, nil
}
