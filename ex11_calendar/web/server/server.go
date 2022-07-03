package server

import (
	"ex11_calendar/model"
	"ex11_calendar/web/config"
	"gopkg.in/yaml.v2"
	"net/http"
	"os"
	"time"
)

type Server struct {
	Config config.Config `yaml:"server"`
	events model.Model   `yaml:"-"`
}

func NewServer(configFile string) (*Server, error) {
	var serv Server

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

func (s *Server) NewRouter() http.Handler {
	return nil
}

func (s *Server) Run() error {

	server := &http.Server{
		Addr:         s.Config.Host + ":" + s.Config.Port,
		Handler:      s.NewRouter(),
		ReadTimeout:  s.Config.Timeout.Read * time.Second,
		WriteTimeout: s.Config.Timeout.Write * time.Second,
		IdleTimeout:  s.Config.Timeout.Idle * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}
