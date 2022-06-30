package main

import (
	"ex10_telnet/pkg"
	"log"
)

func main() {
	t, err := pkg.NewTelnet()
	if err != nil {
		log.Fatalln(err)
	}
	err = t.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
