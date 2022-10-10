package urlshortner

import (
	b64 "encoding/base64"
)

type Store struct {
	Data map[string]string
}

func NewStore() *Store {
	return &Store{Data: map[string]string{}}
}

type Shorter interface {
	Short(string) string
	Find(string) (string, bool)
}

func (s *Store) Short(url string) string {
	shortenURL := s.encode(url)
	s.store(shortenURL, url)
	return shortenURL
}

func (s *Store) Find(url string) (string, bool) {
	if val, ok := s.Data[url]; ok {
		return val, true
	}
	return "", false
}

func (s *Store) store(key, val string) {
	s.Data[key] = val
}

func (s *Store) encode(url string) string {
	return b64.StdEncoding.EncodeToString([]byte(url))
}
