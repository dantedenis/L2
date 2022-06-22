package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
)

const (
	_ = iota
	ErrorScan
	ErrorFprintf
	ErrorQuery
	ErrorPrintln
)

func main() {
	fmt.Println("Get host:")

	var host string
	_, err := fmt.Scan(&host)
	if err != nil {
		os.Exit(ErrorScan)
	}

	time, err := ntp.Time(host)
	if err != nil {
		_, err := fmt.Fprintf(os.Stderr, err.Error())
		if err != nil {
			os.Exit(ErrorFprintf)
		}
		os.Exit(ErrorQuery)
	}

	_, err = fmt.Println(time)
	if err != nil {
		os.Exit(ErrorPrintln)
	}
}
