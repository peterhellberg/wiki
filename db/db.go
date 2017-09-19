package db

import (
	"os"

	"github.com/boltdb/bolt"
)

// DB represents a Bolt-backed data store.
type DB struct {
	*bolt.DB
}

// Open initializes and opens the database.
func (db *DB) Open(path string, mode os.FileMode) error {
	var err error

	db.DB, err = bolt.Open(path, mode, nil)
	if err != nil {
		return err
	}

	// Create buckets.
	err = db.Update(func(tx *Tx) error {
		if _, err := tx.CreateBucketIfNotExists([]byte("pages")); err != nil {
			return &Error{"pages bucket error", err}
		}

		return nil
	})

	if err != nil {
		db.Close()

		return err
	}

	return nil
}

// View executes a function in the context of a read-only transaction.
func (db *DB) View(fn func(*Tx) error) error {
	return db.DB.View(func(tx *bolt.Tx) error {
		return fn(&Tx{tx})
	})
}

// Update executes a function in the context of a writable transaction.
func (db *DB) Update(fn func(*Tx) error) error {
	return db.DB.Update(func(tx *bolt.Tx) error {
		return fn(&Tx{tx})
	})
}
