package server

import (
	"net/http"

	"github.com/peterhellberg/wiki/db"
)

func (s *Server) edit(w http.ResponseWriter, r *http.Request) {
	s.db.View(func(tx *db.Tx) error {
		p, _ := tx.Page(s.getPageName(r))

		return edit.Execute(w, data{
			"Title": string(p.Name),
			"Path":  "/" + string(p.Name),
			"Text":  string(p.Text),
		})
	})
}
