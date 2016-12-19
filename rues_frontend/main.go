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
	var addr = os.Getenv("RUES_FRONTEND_ADDR")
	if addr == "" {
		addr = ":8081" // default binding address
	}

	router := mux.NewRouter().StrictSlash(true)
	router.Handle("/", FrontHandler(RootHandler))

	log.Printf("%s ", "Started ...")
	log.Fatal(http.ListenAndServe(addr, router))
}
