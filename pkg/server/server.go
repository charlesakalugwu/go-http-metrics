package server

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
)

type Server struct {
	router *chi.Mux

	log         *logrus.Entry
	host        string
	port        int
	baseContext string
}

func NewServer(log *logrus.Entry, host string, port int) *Server {
	s := &Server{
		router:      chi.NewRouter(),
		log:         log,
		host:        host,
		port:        port,
		baseContext: "/",
	}
	return s
}

func (s *Server) Address() string {
	return fmt.Sprintf("%s:%d", s.host, s.port)
}

func (s *Server) Run() {
	s.setupRoutes()
	s.log.Infof("starting server on %s", s.Address())
	s.log.WithError(http.ListenAndServe(s.Address(), s.router)).Warn("Server exited.")
}
