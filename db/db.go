package db

import (
	"database/sql"
	"os"
	"sync"

	"github.com/Karitham/shurl/server"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
)

var o sync.Once
var db Database

type Database struct {
	db *sql.DB
}

// DB returns the global database instance
func DB() server.Store {
	o.Do(func() {
		config, err := pgx.ParseConfig(os.Getenv("DATABASE_URL"))
		if err != nil {
			panic(err)
		}
		config.PreferSimpleProtocol = true
		db = Database{db: stdlib.OpenDB(*config)}
	})

	return db
}

func (db Database) Get(key string) (string, error) {
	var value string
	err := db.db.QueryRow("SELECT value FROM urls WHERE key = $1 LIMIT 1", key).Scan(&value)
	if err != nil {
		return "", err
	}
	return value, nil
}

func (db Database) Set(key, value string) error {
	_, err := db.db.Exec("INSERT INTO urls (key, value) VALUES ($1, $2)", key, value)
	if err != nil {
		return err
	}
	return nil
}
