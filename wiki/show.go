package wiki

import (
	"html/template"
	"net/http"

	"github.com/peterhellberg/wiki/db"
	"github.com/zenazn/goji/web"
)

func (w *Wiki) Show(c web.C, rw http.ResponseWriter, r *http.Request) {
	name := w.getPageName(c.URLParams["name"])

	w.DB().View(func(tx *db.Tx) error {
		p, err := tx.Page(name)

		if err != nil {
			p.TextString(emptyPageString)
		}

		vars := map[string]interface{}{
			"Path": "/" + p.Name + "/edit",
			"Text": BytesAsHTML(p.ParsedText()),
		}

		t := template.Must(template.New("show").Parse(showTpl))
		t.Execute(rw, vars)

		return nil
	})
}

func (w *Wiki) RedirectToShow(c web.C, rw http.ResponseWriter, r *http.Request) {
	http.Redirect(rw, r, "/"+c.URLParams["name"], 302)
}

func BytesAsHTML(b []byte) template.HTML {
	return template.HTML(string(b))
}
