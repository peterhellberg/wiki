package db

import (
	"fmt"
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
	db.DB, err = bolt.Open(path, mode)
	if err != nil {
		return err
	}

	// Create buckets.
	err = db.Update(func(tx *Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("pages"))
		assert(err == nil, "pages bucket error: %s", err)

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

// assert will panic with a given formatted message if the given condition is false.
func assert(condition bool, msg string, v ...interface{}) {
	if !condition {
		panic(fmt.Sprintf("assert failed: "+msg, v...))
	}
}
