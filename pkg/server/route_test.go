package server

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/aniruddha2000/kitly/pkg/urlshortner"
)

func TestRoutes(t *testing.T) {
	tests := []struct {
		name       string
		server     *Server
		method     func(*Server) *http.Response
		want       string
		statusCode int
	}{
		{
			name:   "Create short URL",
			server: helperServer(t),
			method: func(s *Server) *http.Response {
				req := httptest.NewRequest(http.MethodPost, "/short?url=github.com", nil)
				w := httptest.NewRecorder()
				s.Short(w, req)
				return w.Result()
			},
			want:       "http://localhost:8989/Z2l0aHViLmNvbQ==",
			statusCode: http.StatusOK,
		},
		{
			name:   "not make POST request",
			server: helperServer(t),
			method: func(s *Server) *http.Response {
				req := httptest.NewRequest(http.MethodGet, "/short?url=github.com", nil)
				w := httptest.NewRecorder()
				s.Short(w, req)
				return w.Result()
			},
			want:       "Request Invalid! POST Request accepted only",
			statusCode: http.StatusBadRequest,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := test.method(test.server)
			defer res.Body.Close()

			if res.StatusCode != test.statusCode {
				t.Errorf("Want %v, got %d", test.statusCode, res.StatusCode)
			}
			got, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Errorf("expected error to be nil got %v", err)
			}
			if !strings.Contains(string(got), test.want) {
				t.Errorf("expected %v got %v", test.want, string(got))
			}
		})
	}
}

func helperServer(t *testing.T) *Server {
	t.Helper()
	router := http.NewServeMux()
	return &Server{Router: router, Shorter: urlshortner.NewStore()}
}
