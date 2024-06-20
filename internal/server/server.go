package server

import (
	"net/http"
	"time"
)

type Server struct {
	Addr string
}

func NewServer() *http.Server {
	s := &Server{
		Addr: "0.0.0.0:6969",
	}

	server := &http.Server{
		Addr:         s.Addr,
		Handler:      s.RegisterRoutes(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return server
}
