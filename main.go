// Copyright 2014-2017 Peter Hellberg.
// Released under the terms of the MIT license.

// wiki is a tiny wiki using BoltDB and Blackfriday.
//
// Installation
//
// You can install wiki with go get:
//
//     go get -u github.com/peterhellberg/wiki
//
// Usage
//
// You can specify one (optional) parameter -db
//
//     PORT=7272 wiki -db="/tmp/foo.db"
//
package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/peterhellberg/wiki/db"
	"github.com/peterhellberg/wiki/server"
)

var dbFile string

func main() {
	flag.StringVar(&dbFile, "db", "/tmp/wiki.db", "Path to the BoltDB file")

	flag.Parse()

	// Setup the logger used by the server
	logger := log.New(os.Stdout, "", 0)

	// Setup the database used by the server
	db, err := newDB(dbFile)
	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()

	hs := setup(logger, db)

	go graceful(hs, 8*time.Second)

	logger.Println("Listening on http://0.0.0.0" + hs.Addr)
	if err := hs.ListenAndServe(); err != nil {
		logger.Fatal(err)
	}
}

func newDB(fn string) (*db.DB, error) {
	db := &db.DB{}

	if err := db.Open(dbFile, 0600); err != nil {
		return nil, err
	}

	return db, nil
}

func setup(logger *log.Logger, db *db.DB) *http.Server {
	return &http.Server{
		Addr:         addr(),
		Handler:      server.New(logger, db),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  2 * time.Minute,
	}
}

func addr() string {
	if port := os.Getenv("PORT"); port != "" {
		return ":" + port
	}

	return ":7272"
}

func graceful(hs *http.Server, timeout time.Duration) {
	signals := make(chan os.Signal, 1)

	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

	<-signals

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	hs.Shutdown(ctx)
}
