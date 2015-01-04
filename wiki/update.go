package wiki

import (
	"net/http"

	"github.com/goji/param"
	"github.com/zenazn/goji/web"

	"github.com/peterhellberg/wiki/db"
)

type formData struct {
	Text string `param:"text"`
}

// Update is the update endpoint of the Wiki
func (w *Wiki) Update(c web.C, rw http.ResponseWriter, r *http.Request) {
	name := w.getPageName(c.URLParams["name"])

	// Parse the POST body
	r.ParseForm()

	var fd formData
	param.Parse(r.Form, &fd)

	w.DB().Update(func(tx *db.Tx) error {
		p := db.Page{Tx: tx, Name: name}
		p.Text = []byte(fd.Text)

		return p.Save()
	})

	path := "/" + string(name)

	if path == "/root" {
		path = "/"
	}

	http.Redirect(rw, r, path, 302)
}
