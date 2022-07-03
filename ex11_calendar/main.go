package main

import (
	"ex11_calendar/web/server"
	"flag"
	"fmt"
	"log"
)

func main() {
	flag.Parse()
	serv, err := server.NewServer(flag.Arg(0))
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(serv)
}
