package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	list, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}
	for {
		conn, err := list.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)
	}

}

func handleClient(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Println(err)
			return
		}
	}(conn)

	buf := make([]byte, 32)

	for {
		_, err := conn.Write([]byte("Hello, whats your name?\n"))
		if err != nil {
			log.Println(err)
			return
		}

		readLen, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			break
		}
		_, err = conn.Write(append([]byte("Goodbye, "), buf[:readLen]...))
		if err != nil {
			log.Println(err)
			return
		}
	}
}
