package main

import (
	"ex11_calendar/web/server"
	"flag"
	"log"
)

func main() {
	flag.Parse()
	serv, err := server.NewServer(flag.Arg(0))
	if err != nil {
		log.Fatalln(err)
	}
	err = serv.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
