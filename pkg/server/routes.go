package server

import (
	"net/http"

	j "github.com/aniruddha2000/kitly/pkg/json"
)

const msg = "http://localhost:8989/"

func (s *Server) initializeRoutes() {
	s.Router.HandleFunc("/short", s.Short)
}

func (s *Server) Short(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		url := r.Form.Get("url")

		shortenURL := s.Shorter.Short(url)

		j.JSON(w, r, http.StatusOK, msg+shortenURL)
	} else {
		j.JSON(w, r, http.StatusBadRequest, "Request Invalid! POST Request accepted only")
	}
}
