package server

import (
	"net/http"

	j "github.com/aniruddha2000/kitly/pkg/json"
)

func (s *Server) initializeRoutes() {
	s.Router.HandleFunc("/short", s.Short)
	s.Router.HandleFunc("/", s.Redirect)
}

func (s *Server) Short(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		url := r.Form.Get("url")

		shortenURL := s.Shorter.Short(url)

		j.JSON(w, r, http.StatusOK, r.Host+"/"+shortenURL)
	} else {
		j.JSON(w, r, http.StatusBadRequest, "Request Invalid! POST Request accepted only")
	}
}

func (s *Server) Redirect(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		id := r.URL.Path

		id = id[1:]

		newurl, ok := s.Shorter.Find(id)
		if !ok {
			j.JSON(w, r, http.StatusNotFound, "Opps something wrong!")
		}

		newurl = "http://" + newurl + "/"

		http.Redirect(w, r, newurl, http.StatusPermanentRedirect)
	} else {
		j.JSON(w, r, http.StatusBadRequest, "Request Invalid! GET Request accepted only")
	}
}
