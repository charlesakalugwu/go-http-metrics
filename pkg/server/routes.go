package server

import (
	"github.com/go-chi/chi/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func (s *Server) setupRoutes() {
	s.router.Use(middleware.DefaultCompress)
	s.router.Use(middleware.RedirectSlashes)
	s.router.Use(middleware.Recoverer)

	s.router.Use(s.requestInfo)

	s.router.Get("/", handleGet)
	s.router.Get("/hello", handleHelloWorld)
	s.router.Handle("/metrics", promhttp.Handler())
}
