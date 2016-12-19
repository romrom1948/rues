// Copyright (C) 2016 romrom@tutanota.com
// Use of this source code is governed by the GPLv3
// license that can be found in the LICENSE file.

package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type func(w http.ResponseWriter, r *http.Request, db *sql.DB) error

func (h DBHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var e = GetBackEnv()
	start := time.Now()
	e := h.handler(w, r, e.db)
	if e != nil {
		http.Error(w, `500 internal server error`, http.StatusInternalServerError)
		log.Printf("%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			e,
		)
		return
	} else {
		log.Printf("%s\t%s\t(%s)",
			r.Method,
			r.RequestURI,
			time.Since(start),
		)
	}
}

func JsonHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
}
