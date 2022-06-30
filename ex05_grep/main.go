package main

import (
	"ex05_grep/grep"
	"log"
)

func main() {
	g := grep.NewGrep()
	log.Println(g)
	g.Run()
	g.PrintMap()
}
