package main

import (
	"flag"
	"log"

	"github.com/peterhellberg/wiki/db"
	"github.com/peterhellberg/wiki/wiki"
	"github.com/zenazn/goji"
)

var dbFile = flag.String("db", "/tmp/wiki.db", "Path to the BoltDB file")

func main() {
	flag.Parse()

	// Initialize db.
	var db db.DB
	if err := db.Open(*dbFile, 0600); err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Initialize wiki.
	w, err := wiki.NewWiki(&db)
	if err != nil {
		log.Fatal(err)
	}

	// Setup up the routes for the wiki
	goji.Get("/", w.Show)
	goji.Get("/:name", w.Show)
	goji.Get("/:name/", w.RedirectToShow)
	goji.Get("/:name/edit", w.Edit)
	goji.Post("/:name", w.Update)

	// Start the web server
	goji.Serve()
}
