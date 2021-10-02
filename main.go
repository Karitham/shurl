package main

import (
	"encoding/base64"
	"log"
	"net/http"

	"github.com/Karitham/shurl/db"
	"github.com/Karitham/shurl/keys"
	"github.com/Karitham/shurl/server"
)

func main() {
	// New Key generator, 4 byte keys
	// 2^32 => 4.294.967.296 possible keys,
	// clash is unlikely at small scale, but possible.
	gen := keys.New(4)

	// Encoder
	coder := base64.RawURLEncoding

	store, cancel, err := db.New("shurl.db")
	if err != nil {
		log.Fatal("error initializing the store", err)
	}
	defer cancel()

	h := server.NewServer("/shurl", store, gen, coder)

	log.Println("serving on :8080")
	if err := http.ListenAndServe(":8080", h); err != nil {
		log.Println("error serving, exiting", err)
	}
}
