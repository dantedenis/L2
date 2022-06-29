package main

import (
	"ex06_cut/cut"
	"fmt"
)

func main() {
	c := cut.NewCut()
	fmt.Println(c.String())
	err := c.Run()
	if err != nil {
		return
	}
}
