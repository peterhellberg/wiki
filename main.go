// Copyright 2014-2015 Peter Hellberg.
// Released under the terms of the MIT license.

// wiki is a tiny wiki using Goji, BoltDB and Blackfriday.
//
// Installation
//
//     go get -u github.com/peterhellberg/wiki
//
// Usage
//
// You can specify two (optional) parameters -bind and -db
//
//     wiki -bind=":7272" -db="/tmp/foo.db"
//
package main

import (
	"flag"
	"log"

	"github.com/peterhellberg/wiki/db"
	"github.com/peterhellberg/wiki/wiki"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web/middleware"
)

var dbFile = flag.String("db", "/tmp/wiki.db", "Path to the BoltDB file")
var loggerEnabled = flag.Bool("logger-enabled", true, "Enable/Disable logging to stdout")

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

	if *loggerEnabled != true {
		goji.Abandon(middleware.Logger)
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
