package db

import (
	"github.com/boltdb/bolt"
)

// Tx represents a BoltDB transaction
type Tx struct {
	*bolt.Tx
}

// Page retrieves a Page from the database with the given name.
func (tx *Tx) Page(name []byte) (*Page, error) {
	p := &Page{Tx: tx, Name: name}
	return p, p.Load()
}
