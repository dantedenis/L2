package server

import (
	"ex11_calendar/model"
	"ex11_calendar/web/config"
	"gopkg.in/yaml.v2"
	"log"
	"net/http"
	"os"
	"time"
)

type Server struct {
	Config config.Config      `yaml:"server"`
	Events model.EventManager `yaml:"-"`
}

func NewServer(configFile string) (*Server, error) {
	var serv Server
	_ = configFile
	serv.Events = *model.NewEventManager()

	file, err := os.Open(configFile)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err = file.Close()
	}(file)

	if err = yaml.NewDecoder(file).Decode(&serv); err != nil {
		return nil, err
	}

	return &serv, nil
}

func (s *Server) Run() error {

	server := &http.Server{
		Addr:         s.Config.Host + ":" + s.Config.Port,
		Handler:      s.NewRouter(),
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	log.Println("Run server, host:", s.Config.Host, "port:", s.Config.Port)
	err := server.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}
