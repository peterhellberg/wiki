package wiki

import "github.com/peterhellberg/wiki/db"

type Wiki struct {
	db *db.DB
}

func NewWiki(db *db.DB) (*Wiki, error) {
	// Setup the wiki.
	w := &Wiki{db: db}

	return w, nil
}

// DB returns the database associated with the handler.
func (w *Wiki) DB() *db.DB {
	return w.db
}

func (w *Wiki) getPageName(name string) string {
	if name == "" {
		return "root"
	}

	return name
}
