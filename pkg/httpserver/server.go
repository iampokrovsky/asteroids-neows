package httpserver

import (
	"net"
	"net/http"
	"time"
)

const (
	defaultReadTimeout  = 5 * time.Second
	defaultWriteTimeout = 5 * time.Second
	defaultAddr         = ":8080"
)

type Server struct {
	server *http.Server
}

func New(handler http.Handler, port string) *Server {
	addr := defaultAddr

	if port != "" {
		addr = net.JoinHostPort("", port)
	}

	httpServer := &http.Server{
		Handler:      handler,
		ReadTimeout:  defaultReadTimeout,
		WriteTimeout: defaultWriteTimeout,
		Addr:         addr,
	}

	s := &Server{
		server: httpServer,
	}

	return s
}

func (s *Server) Run() error {
	return s.server.ListenAndServe()
}
