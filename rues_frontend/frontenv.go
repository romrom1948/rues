// Copyright (C) 2016 romrom@tutanota.com
// Use of this source code is governed by the GPLv3
// license that can be found in the LICENSE file.

package main

import (
	"html/template"
	"sync"
	"fmt"
	"os"
)

type FrontEnv struct {
	Backend_URL string
	Root *template.Template
}

var env *FrontEnv
var once sync.Once

func GetFrontEnv() *FrontEnv {
	once.Do(func() {
		tmpl, err := template.New("rues.html").ParseFiles(os.Getenv("RUES_TEMPLATE_ROOT"))

		if err != nil {
			fmt.Println("error parsing template !")
			os.Exit(-1)
		}
		
		env = &FrontEnv{
			Backend_URL: os.Getenv("RUES_BACKEND_ADDR"),
			Root: tmpl}
	})
	return env
}
