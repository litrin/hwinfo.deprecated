package main

import (
	"encoding/json"
	"fmt"
	"log"

	//	"github.com/mickep76/hwinfo/memory"
	//	"github.com/mickep76/hwinfo/system"
	//	"github.com/mickep76/hwinfo/cpu"
	//	"github.com/mickep76/hwinfo/network"
	"github.com/mickep76/hwinfo/opsys"
)

func main() {
	d := opsys.New()
	err := d.Get()
	if err != nil {
		log.Fatal(err.Error())
	}

	b, err2 := json.MarshalIndent(d, "", "    ")
	if err2 != nil {
		fmt.Println(err2.Error())
	}

	fmt.Println(string(b))
}
