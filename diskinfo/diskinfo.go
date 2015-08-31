package diskinfo

// Info structure for information about a systems memory.
type Info struct {
	Disks []Disk `json:"disk"`
}

type Disk struct {
	Device string `json:"device"`
	Name   string `json:"name"`
	//	Major  int    `json:"major"`
	//	Minor  int    `json:"minor"`
	//	Blocks int    `json:"blocks"`
	SizeGB int `json:"size_gb"`
}

// PVs

type VolGrp struct {
	Name string `json:"name"`
	PV
}

/*
$ vgs --separator ,

  VG,#PV,#LV,#SN,Attr,VSize,VFree
  vg_fs,1,1,0,wz--n-,278.46g,0
  vg_root,1,4,0,wz--n-,278.00g,104.41g
*/
