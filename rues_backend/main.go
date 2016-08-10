// Copyright (C) 2016 romrom@tutanota.com
// Use of this source code is governed by the GPLv3
// license that can be found in the LICENSE file.

package main

import (
	"os"
	"strings"
	"fmt"
    "log"
    "net/http"
    
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"	
	
	. "github.com/romrom1948/rues/util"	
)

var helpMessage = []string{
	"Usage: rues_backend <db> <addr>",
	"Start a JSON backend server for rues db <db>.",
	"<addr> is optional. The server will bind on it if supplied.",
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
    
	router.Handle("/communes", DBHandler(CommunesHandler))
	router.Handle("/commune/name/{commune}", DBHandler(CommuneNameHandler))
	router.Handle("/commune/id/{id}", DBHandler(CommuneIdHandler))
	router.Handle("/commune/like/{request}", DBHandler(CommuneLikeHandler))
	
    router.Handle("/voies", DBHandler(VoiesHandler))
	router.Handle("/voie/name/{voie}", DBHandler(VoieNameHandler))
	router.Handle("/voie/id/{id}", DBHandler(VoieIdHandler))
	router.Handle("/voie/like/{request}", DBHandler(VoieLikeHandler))
	
	log.Printf("%s ", "Started ...");
    log.Fatal(http.ListenAndServe(addr, router))	
}
