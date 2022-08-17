package kvdb

import (
	"errors"

	bolt "go.etcd.io/bbolt"
)

type Store struct {
	db        *bolt.DB
	filepath  string
	urlBucket string
}

// New returns a new bbolt store and a function to close it.
func New(filepath string) (s *Store, cancel func(), err error) {
	s = &Store{
		filepath:  filepath,
		urlBucket: "URLS",
	}

	db, err := bolt.Open(filepath, 0666, nil)
	if err != nil {
		return s, nil, err
	}

	tx, err := db.Begin(true)
	if err != nil {
		return nil, nil, err
	}

	_, err = tx.CreateBucketIfNotExists([]byte(s.urlBucket))
	if err != nil {
		return nil, nil, err
	}

	tx.Commit()

	s.db = db

	return s,
		func() { db.Close() },
		nil
}

func (s *Store) Get(key []byte) ([]byte, error) {
	var value []byte
	err := s.db.View(func(tx *bolt.Tx) error {
		value = tx.Bucket([]byte(s.urlBucket)).Get(key)
		if value == nil || len(value) < 3 {
			return errors.New("value not found")
		}

		return nil
	})
	if err != nil {
		return nil, err
	}
	return value, nil
}

func (s *Store) Set(key, value []byte) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		err := tx.Bucket([]byte(s.urlBucket)).Put(key, value)
		if err != nil {
			return err
		}
		return nil
	})
}
