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
	// считывает хоста с инпута
	_, err := fmt.Scan(&host)
	if err != nil {
		os.Exit(ErrorScan)
	}

	// и отправляем хоста в бибилотеку для коннекта с сервером и запроса времени
	time, err := ntp.Time(host)
	if err != nil {
		_, err := fmt.Fprintf(os.Stderr, err.Error())
		if err != nil {
			os.Exit(ErrorFprintf)
		}
		os.Exit(ErrorQuery)
	}
	// проверяем на ошибки, если все ок печатаем результат
	_, err = fmt.Println(time)
	if err != nil {
		os.Exit(ErrorPrintln)
	}
}
