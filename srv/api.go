package srv

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Store interface {
	Get(key []byte) ([]byte, error)
	Set(key, value []byte) error
}

type Generator interface {
	Generate() ([]byte, error)
}

type Coder interface {
	DecodeString(string) ([]byte, error)
	EncodeToString([]byte) string
}

func NewServer(prefix string, s Store, g Generator, c Coder) http.Handler {
	r := chi.NewMux()
	r.Route(prefix, func(r chi.Router) {
		r.Get("/{key}", GetKey(s, c))
		r.Put("/", SetKey(s, g, c))
	})
	return r
}

func GetKey(s Store, c Coder) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := chi.URLParam(r, "key")
		if len(key) == 0 {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		decoded, err := c.DecodeString(key)
		if err != nil {
			http.Error(w, "Error decoding Key", http.StatusInternalServerError)
			return
		}

		val, err := s.Get(decoded)
		if err != nil {
			http.Error(w, "Error getting key", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, string(val), http.StatusMovedPermanently)
	}
}

func SetKey(s Store, g Generator, c Coder) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u := r.URL.Query().Get("url")

		if len(u) == 0 {
			http.Error(w, "Not Found", http.StatusNotFound)
			return
		}

		key, err := g.Generate()
		if err != nil {
			http.Error(w, "Error getting key", http.StatusInternalServerError)
			return
		}

		if err := s.Set(key, []byte(u)); err != nil {
			http.Error(w, "Error getting key", http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "%s\n", c.EncodeToString(key))
	}
}
