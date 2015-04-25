package wiki

import (
	"net/http"
	"text/template"

	"github.com/zenazn/goji/web"

	"github.com/peterhellberg/wiki/db"
)

// Edit is the edit endpoint of the Wiki
func (w *Wiki) Edit(c web.C, rw http.ResponseWriter, r *http.Request) {
	name := w.getPageName(c.URLParams["name"])

	w.DB().View(func(tx *db.Tx) error {
		p, _ := tx.Page(name)

		t, err := template.New("edit").Parse(editTpl)

		template.Must(t, err).Execute(rw, map[string]string{
			"Path": "/" + string(p.Name),
			"Text": string(p.Text),
		})

		return nil
	})
}
