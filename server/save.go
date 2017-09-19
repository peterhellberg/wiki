package server

import (
	"net/http"
	"strings"

	"github.com/peterhellberg/wiki/db"
)

func (s *Server) save(w http.ResponseWriter, r *http.Request) {
	s.db.Update(func(tx *db.Tx) error {
		r.ParseForm()

		p := db.Page{
			Tx:   tx,
			Name: s.getPageName(r),
		}

		p.Text = []byte(strings.TrimSpace(r.FormValue("text")))

		return p.Save()
	})

	s.redirect(w, r)
}
