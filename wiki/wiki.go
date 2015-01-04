package wiki

import "github.com/peterhellberg/wiki/db"

// Wiki represents the entire Wiki, contains the db
type Wiki struct {
	db *db.DB
}

// NewWiki creates a new Wiki
func NewWiki(db *db.DB) (*Wiki, error) {
	// Setup the wiki.
	w := &Wiki{db: db}

	return w, nil
}

// DB returns the database associated with the handler.
func (w *Wiki) DB() *db.DB {
	return w.db
}

func (w *Wiki) getPageName(name string) []byte {
	if name == "" {
		return []byte("root")
	}

	return []byte(name)
}
