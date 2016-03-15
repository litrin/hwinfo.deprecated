package main

import (
	"encoding/json"
	"fmt"
	"log"

	//	"github.com/mickep76/hwinfo/system"
	"github.com/mickep76/hwinfo/cpu"
)

func main() {
	e := cpu.New()
	if err := e.Get(); err != nil {
		log.Fatal(err.Error())
	}

	data, err := json.MarshalIndent(e.Data(), "", "    ")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(data))

	cache, err := json.MarshalIndent(e.Cache(), "", "    ")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(cache))
}
