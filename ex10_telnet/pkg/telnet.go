package pkg

import (
	"errors"
	"flag"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"time"
)

type telnet struct {
	timeout    time.Duration
	host, port string
}

// NewTelnet Парсим аргументы, проверяем их количество и заполняем структуру
func NewTelnet() (*telnet, error) {
	var t time.Duration
	flag.DurationVar(&t, "timeout", 10*time.Second, "timeout")
	flag.Parse()

	args := flag.Args()
	if len(args) != 2 {
		return nil, errors.New("Invalid number of arguments")
	}

	return &telnet{
		timeout: t,
		host:    args[0],
		port:    args[1],
	}, nil
}

func (t *telnet) Run() error {

	errChan, sign := make(chan error), make(chan os.Signal)

	// Создаем и подключаемся по ТСР с имеющими у нас хостом и портом + таймаут ожидания
	conn, err := net.DialTimeout("tcp", t.host+":"+t.port, t.timeout)
	if err != nil {
		return err
	}

	defer func() {
		//отмечаемся, что соединение было закрыто
		log.Println("Close connection")
		err = conn.Close()
	}()

	// биндим обработчик сигнала
	signal.Notify(sign, os.Interrupt)

	// socket -> stdout
	go func() {
		// отправляет байты из одного дескриптора в другой
		_, err := io.Copy(os.Stdout, conn)
		errChan <- err
	}()

	// stdin -> socket
	go func() {
		// отправляет байты из одного дескриптора в другой
		_, err := io.Copy(conn, os.Stdin)
		errChan <- err
	}()

	// Ожидание тригерра на каналах и реакция на них
	select {
	case c := <-sign:
		log.Println("Catch signal:", c)
		return nil
	case err := <-errChan:
		return err
	}
}
