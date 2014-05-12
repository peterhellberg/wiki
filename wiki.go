package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/boltdb/bolt"
	"github.com/russross/blackfriday"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/param"
	"github.com/zenazn/goji/web"
)

func main() {
	// Routes for the wiki
	goji.Get("/:name", show)
	goji.Get("/:name/edit", edit)
	goji.Post("/:name", update)
	goji.Get("/", show)

	// Start the server
	goji.Serve()
}

func show(c web.C, w http.ResponseWriter, r *http.Request) {
	path := "foo.db"
	bucket := "wiki"
	name := c.URLParams["name"]

	if name == "" {
		name = "root"
	}

	// Open database
	db, err := bolt.Open(path, 0600)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer db.Close()

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))

		if b == nil {
			return fmt.Errorf("bucket not found: %s", bucket)
		}

		// Get value
		val := b.Get([]byte(name))
		if val == nil {
			val = []byte{}
		}

		p := Page{Text: string(val)}

		vars := map[string]interface{}{
			"Path": "/" + name + "/edit",
			"Text": template.HTML(string(p.ParsedText())),
		}

		t := template.Must(template.New("page").Parse(page))
		t.Execute(w, vars)

		return nil
	})
}

type EditForm struct {
	Path string
	Text string
}

func edit(c web.C, w http.ResponseWriter, r *http.Request) {
	path := "foo.db"
	bucket := "wiki"
	name := c.URLParams["name"]

	// Open database
	db, err := bolt.Open(path, 0600)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer db.Close()

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))

		if b == nil {
			return fmt.Errorf("bucket not found: %s", bucket)
		}

		// Get value
		val := b.Get([]byte(name))
		if val == nil {
			val = []byte{}
		}

		ef := EditForm{
			Path: "/" + name,
			Text: string(val),
		}

		f := template.Must(template.New("editForm").Parse(editForm))
		f.Execute(w, ef)

		return nil
	})
}

type Page struct {
	Text string `param:"text"`
}

func (p *Page) ParsedText() []byte {
	return blackfriday.MarkdownCommon([]byte(p.Text))
}

func update(c web.C, w http.ResponseWriter, r *http.Request) {
	name := c.URLParams["name"]

	var page Page

	// Parse the POST body
	r.ParseForm()
	err := param.Parse(r.Form, &page)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	Put("foo.db", "wiki", name, []byte(page.Text))

	http.Redirect(w, r, "/"+name, http.StatusMovedPermanently)
}

const page = `<!DOCTYPE html>
<html>
	<head>
	</head>
	<body style="padding: 0.5em;">
	<a href="{{.Path}}">edit</a>
	<div class="container" style="padding: 1em; border: 1px solid #ccc;">{{.Text}}</div>
	</body>
</html>`

const editForm = `<!DOCTYPE html>
<html>
	<head>
	</head>
	<body>
		<div class="container">
			<form action="{{.Path}}" method="POST">
				<textarea name="text" cols="80" rows="25">{{.Text}}</textarea>
				<input type="submit" value="Update" />
			</form>
		</div>
	</body>
</html>`
