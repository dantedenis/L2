package main

import (
	"ex08_community_os/execute"
	"log"
)

func main() {
	shell, err := execute.NewShell()
	if err != nil {
		log.Println(err)
		return
	}
	err = shell.Run()
	if err != nil {
		log.Println(err)
		return
	}
}
