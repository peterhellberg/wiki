package wiki

import (
	"net/http"
	"text/template"

	"github.com/peterhellberg/wiki/db"
	"github.com/zenazn/goji/web"
)

func (w *Wiki) Edit(c web.C, rw http.ResponseWriter, r *http.Request) {
	name := w.getPageName(c.URLParams["name"])

	w.DB().View(func(tx *db.Tx) error {
		p, _ := tx.Page(name)

		t, err := template.New("edit").Parse(editTpl)

		template.Must(t, err).Execute(rw, map[string]string{
			"Path": "/" + p.Name,
			"Text": string(p.Text),
		})

		return nil
	})
}
