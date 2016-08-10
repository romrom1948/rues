// Copyright (C) 2016 romrom@tutanota.com
// Use of this source code is governed by the GPLv3
// license that can be found in the LICENSE file.

package util

import (
	"time"
    "log"
    "net/http"
	"database/sql"
	"os"
	
	_ "github.com/mattn/go-sqlite3"	
)

// custom type for db-using handler
type DBHandler func(w http.ResponseWriter, r *http.Request, db *sql.DB) (error)
func (h DBHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", os.Args[1])
	if err != nil {
		http.Error(w, `500 internal server error`, http.StatusInternalServerError) 
		log.Printf("%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			err,
		)
		return		
	}	
	defer db.Close()

	start := time.Now()
	e := h(w, r, db)
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

// default header for JSON backend responses
func JsonHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "GET")
}
