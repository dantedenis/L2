package main

import (
	"ex05_grep/grep"
	"fmt"
	"os"
)

// for run
// cat main.go | go run main.go {option&patter}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("--help for more info")
		return
	}
	g := grep.NewGrep()
	g.Run()
}
