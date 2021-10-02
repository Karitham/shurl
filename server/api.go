package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

// mocks with https://github.com/derision-test/go-mockgen

//go:generate go-mockgen -f github.com/Karitham/shurl/server -i Store -o store_mock.go
type Store interface {
	Get(key []byte) ([]byte, error)
	Set(key, value []byte) error
}

//go:generate go-mockgen -f github.com/Karitham/shurl/server -i Generator -o generator_mock.go
type Generator interface {
	Generate() ([]byte, error)
}

//go:generate go-mockgen -f github.com/Karitham/shurl/server -i Coder -o coder_mock.go
type Coder interface {
	DecodeString(string) ([]byte, error)
	EncodeToString([]byte) string
}

func NewServer(prefix string, s Store, g Generator, c Coder) http.Handler {
	srv := &Server{
		store: s,
		gen:   g,
		coder: c,
	}

	r := chi.NewMux()
	r.Route(prefix, func(r chi.Router) {
		r.Get("/{key}", srv.redirectFromURL)
		r.Put("/", srv.newShortURL)
	})

	return r
}

type Server struct {
	store Store
	gen   Generator
	coder Coder
}

func (s *Server) redirectFromURL(w http.ResponseWriter, r *http.Request) {
	key := chi.URLParam(r, "key")
	if len(key) == 0 {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	decoded, err := s.coder.DecodeString(key)
	if err != nil {
		http.Error(w, "Error decoding Key", http.StatusInternalServerError)
		return
	}

	val, err := s.store.Get(decoded)
	if err != nil {
		http.Error(w, "Error getting key", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, string(val), http.StatusMovedPermanently)
}

func (s *Server) newShortURL(w http.ResponseWriter, r *http.Request) {
	u := r.URL.Query().Get("url")

	if len(u) == 0 {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	key, err := s.gen.Generate()
	if err != nil {
		http.Error(w, "Error getting key", http.StatusInternalServerError)
		return
	}

	if err := s.store.Set(key, []byte(u)); err != nil {
		http.Error(w, "Error getting key", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(s.coder.EncodeToString(key)))
}
