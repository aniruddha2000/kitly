package server

import (
	"log"
	"net/http"
	"os"

	"github.com/aniruddha2000/kitly/pkg/urlshortner"
)

type Server struct {
	Router  *http.ServeMux
	Shorter urlshortner.Shorter
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Initialize() {
	s.Router = http.NewServeMux()

	f, err := os.Open("db.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s.Shorter = urlshortner.NewFileStore(f)
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
