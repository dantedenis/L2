package main

import (
	"ex06_cut/cut"
	"log"
)

func main() {
	c := cut.NewCut()
	err := c.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
