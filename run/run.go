package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/mickep76/hwinfo"
)

func main() {
	d := hwinfo.NewCached()
	//	d := hwinfo.New()
	if err := d.Get(); err != nil {
		log.Fatal(err.Error())
	}

	b, err := json.MarshalIndent(d, "", "    ")
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(string(b))
}
