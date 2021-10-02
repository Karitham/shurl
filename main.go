package main

import (
	"encoding/base32"
	"log"
	"net/http"

	"github.com/Karitham/shurl/db"
	"github.com/Karitham/shurl/keys"
	"github.com/Karitham/shurl/srv"
)

func main() {
	// New Key generator, 64 byte keys
	g := keys.New(3)
	// Encoder
	coder := base32.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZ234567").WithPadding(base32.NoPadding)

	store, cancel, err := db.New("shurl.db")
	if err != nil {
		log.Fatal("error initializing the store", err)
	}
	defer cancel()

	h := srv.NewServer("/shurl", store, g, coder)

	log.Println("serving on :8080")

	if err := http.ListenAndServe(":8080", h); err != nil {
		log.Println("error serving, exiting", err)
	}
}
