// Copyright (C) 2016 romrom@tutanota.com
// Use of this source code is governed by the GPLv3
// license that can be found in the LICENSE file.

package main

import (
	"os"
	"strings"
	"fmt"
    "log"
    "time"
    "net/http"
    "database/sql"
    
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"	
	
	"github.com/romrom1948/rues"
)

var helpMessage = []string{
	"Usage: rues_backend <db> <port>",
	"Start a JSON backend server for rues db <db>.",
	"<port> is optional. The server will bind on it if supplied.",
	"",
}

// custom handler type for DB management, headers setting and logging
type handler func(w http.ResponseWriter, r *http.Request, db *sql.DB)
func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	
	start := time.Now()
	
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "GET")
    w.WriteHeader(http.StatusOK)
	
	h(w, r, db)
	
	log.Printf("%s\t%s\t(%s)",
		r.Method,
		r.RequestURI,
		time.Since(start),
	)
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println(strings.Join(helpMessage, "\n"))
		fmt.Println("need an sqlite3 database path !")
		os.Exit(-1)
	}
	
	var addr = ":8080" // default binding address
	if len(os.Args) == 3 {
		addr = os.Args[2]
	}
	
    router := mux.NewRouter().StrictSlash(true)
	// routes are set below, not enough of them to warrant a specific file
    
	router.Handle("/communes", handler(rues.CommunesHandler))
	router.Handle("/commune/name/{commune}", handler(rues.CommuneNameHandler))
	router.Handle("/commune/id/{id}", handler(rues.CommuneIdHandler))
	router.Handle("/commune/like/{request}", handler(rues.CommuneLikeHandler))	
	
    router.Handle("/voies", handler(rues.VoiesHandler))
	router.Handle("/voie/name/{voie}", handler(rues.VoieNameHandler))
	router.Handle("/voie/id/{id}", handler(rues.VoieIdHandler))
	router.Handle("/voie/like/{request}", handler(rues.VoieLikeHandler))
	
	log.Printf("%s ", "Started ...");
    log.Fatal(http.ListenAndServe(addr, router))	
}
