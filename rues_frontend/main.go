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
    
	"github.com/gorilla/mux"
	
	"github.com/romrom1948/rues"
)

var helpMessage = []string{
	"Usage: rues_frontend <backend_addr> <addr>",
	"Start a JSON frontend server for a backend reachable at <backend_addr.",
	"<addr> is optional. The server will bind on it if supplied.",
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println(strings.Join(helpMessage, "\n"))
		fmt.Println("need a backend adress !")
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
