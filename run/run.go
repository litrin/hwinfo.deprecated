package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/mickep76/hwinfo/memory"
)

func main() {
	d := memory.New()
	d.Get()

	b, err := json.MarshalIndent(d, "", "    ")
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(string(b))

	time.Sleep(2 * time.Second)

	d.Get()

	time.Sleep(30 * time.Second)

	d.Get()

	b, err = json.MarshalIndent(d, "", "    ")
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(string(b))
}
