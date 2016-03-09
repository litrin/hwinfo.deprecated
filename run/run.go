package main

import (
	"encoding/json"
	"fmt"

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
}
