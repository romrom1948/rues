// Copyright (C) 2016 romrom@tutanota.com
// Use of this source code is governed by the GPLv3
// license that can be found in the LICENSE file.

package main

import (
	"database/sql"
	"log"
	"os"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

type BackEnv struct {
	db *sql.DB
}

var env *BackEnv
var once sync.Once

func GetBackEnv() *BackEnv {
	once.Do(func() {
		var db_path = os.Getenv("RUES_DB")
		if db_path == "" {
			db_path = "rues.db" // default sql database path
		}

		// check for existence first, since sql.Open() will open an empty file without error
		if _, err := os.Stat(db_path); os.IsNotExist(err) {
			log.Printf("db file %s does not exist !", db_path)
			os.Exit(-1)
		}
		db, err := sql.Open("sqlite3", db_path)
		if err != nil {
			log.Printf("unable to open db file %s: %s !", db_path, err)
			os.Exit(-1)
		}

		env = &BackEnv{db}
	})
	return env
}
