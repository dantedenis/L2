package main

import (
	"ex03_sort/pkg"
	"log"
)

func main() {
	sort, err := pkg.NewSort()
	if err != nil {
		log.Fatalln(err)
		return
	}
	sort.Run()
	_, err = sort.Write()
	if err != nil {
		log.Fatalln(err)
		return
	}
}
