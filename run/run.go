package main

import (
	"encoding/json"
	"fmt"
	"log"

	//	"github.com/mickep76/hwinfo/memory"
	//	"github.com/mickep76/hwinfo/system"
	//	"github.com/mickep76/hwinfo/network"
	//	"github.com/mickep76/hwinfo/opsys"
	//	"github.com/mickep76/hwinfo/pci"
	//	"github.com/mickep76/hwinfo/sysctl"
	//	"github.com/mickep76/hwinfo/disk"
	"github.com/mickep76/hwinfo/dock2box"
	//	"github.com/mickep76/hwinfo/lvm"
	//	"github.com/mickep76/hwinfo/mount"
	//	"github.com/mickep76/hwinfo/cpu"
	//	"github.com/mickep76/hwinfo/routes"
)

func main() {
	d := dock2box.NewCached()
	if err := d.Get(); err != nil {
		log.Fatal(err.Error())
	}

	b, err := json.MarshalIndent(d, "", "    ")
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(string(b))
}
