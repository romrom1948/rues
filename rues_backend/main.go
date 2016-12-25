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

	router.Handle("/communes", DBHandler(CommunesHandler))
	router.Handle("/commune/name/{commune}", DBHandler(CommuneNameHandler))
	router.Handle("/commune/id/{id}", DBHandler(CommuneIdHandler))
	router.Handle("/commune/like/{request}", DBHandler(CommuneLikeHandler))

	router.Handle("/voies", DBHandler(VoiesHandler))
	router.Handle("/voie/name/{voie}", DBHandler(VoieNameHandler))
	router.Handle("/voie/id/{id}", DBHandler(VoieIdHandler))
	router.Handle("/voie/like/{request}", DBHandler(VoieLikeHandler))

	log.Printf("%s ", "Started ...")
	log.Fatal(http.ListenAndServe(addr, router))
}
