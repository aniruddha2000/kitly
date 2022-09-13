package server

import (
	"log"
	"net/http"
)

type Server struct {
	Router *http.ServeMux
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Initialize() {
	s.Router = http.NewServeMux()
	s.initializeRoutes()
}

func (s *Server) Run(addr string) {
	server := &http.Server{
		Addr:    ":" + addr,
		Handler: s.Router,
	}

	log.Printf("Server lsitenning to port %s\n", addr)
	log.Fatal(server.ListenAndServe())
}
