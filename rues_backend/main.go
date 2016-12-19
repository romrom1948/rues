// Copyright (C) 2016 romrom@tutanota.com
// Use of this source code is governed by the GPLv3
// license that can be found in the LICENSE file.

package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	var addr = os.Getenv("RUES_BACKEND_ADDR")
	if addr == "" {
		addr = ":8080" // default binding address
	}

	router := mux.NewRouter().StrictSlash(true)
	// routes are set below, not enough of them to warrant a specific file

	router.Handle("/communes", DBHandler{db, CommunesHandler})
	router.Handle("/commune/name/{commune}", DBHandler{db, CommuneNameHandler})
	router.Handle("/commune/id/{id}", DBHandler{db, CommuneIdHandler})
	router.Handle("/commune/like/{request}", DBHandler{db, CommuneLikeHandler})

	router.Handle("/voies", DBHandler{db, VoiesHandler})
	router.Handle("/voie/name/{voie}", DBHandler{db, VoieNameHandler})
	router.Handle("/voie/id/{id}", DBHandler{db, VoieIdHandler})
	router.Handle("/voie/like/{request}", DBHandler{db, VoieLikeHandler})

	log.Printf("%s ", "Started ...")
	log.Fatal(http.ListenAndServe(addr, router))
}
