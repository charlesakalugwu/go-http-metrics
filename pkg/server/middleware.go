package server

import "net/http"

func (s *Server) requestInfo(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.log.Debugf("starting: %s %s", r.Method, r.URL)
		handler.ServeHTTP(w, r)
		s.log.Debugf("ending:   %s %s", r.Method, r.URL)
	})
}
