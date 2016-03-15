package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/mickep76/hwinfo"
)

func main() {
	i := hwinfo.New()
	if err := i.Update(); err != nil {
		log.Fatal(err.Error())
	}

	data, err := json.MarshalIndent(i.GetData(), "", "    ")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(data))

	cache, err := json.MarshalIndent(i.GetCache(), "", "    ")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(cache))
}
