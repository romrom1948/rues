// Copyright (C) 2016 romrom@tutanota.com
// Use of this source code is governed by the GPLv3
// license that can be found in the LICENSE file.

package main

import (
	"net/http"
	"time"
	"log"
)

type FrontHandler func(w http.ResponseWriter, r *http.Request, env *FrontEnv) error

func (h FrontHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	e := h(w, r, GetFrontEnv())
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

func FrontHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.Header().Set("Content-Language", "fr")
}

func RootHandler(w http.ResponseWriter, r *http.Request, env *FrontEnv) error {
	return env.Root.Execute(w, env)
}
