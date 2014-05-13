package db

import (
	"github.com/boltdb/bolt"
)

type Tx struct {
	*bolt.Tx
}

// Page retrieves a Page from the database with the given name.
func (tx *Tx) Page(name string) (*Page, error) {
	p := &Page{Tx: tx, Name: name}
	return p, p.Load()
}
