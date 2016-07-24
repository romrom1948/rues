// Copyright (C) 2016 romrom@tutanota.com
// Use of this source code is governed by the GPLv3
// license that can be found in the LICENSE file.

package rues

import (
	"fmt"
    "log"
    "net/http"
	"database/sql"
	"encoding/json"
	
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	
	"github.com/romrom1948/rues/util"	
)

func VoiesHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	rows, err := db.Query("SELECT id, nom, occurences FROM voies")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var voies util.Voies
	for rows.Next() {
		var voie util.Voie

		err = rows.Scan(&voie.Id, &voie.Nom, &voie.Occurences)
		if err != nil {
			log.Fatal(err)
		} 
	
		voies = append(voies, voie)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(voies)	
}

func VoieNameHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	vars := mux.Vars(r)
	voie := vars["voie"]	

	rows, err := db.Query(`
						  SELECT communes.id, communes.nom, communes.cp, communes.voies 
							FROM voies 
							INNER JOIN liens ON liens.id_voie=voies.id 
							INNER JOIN communes ON liens.id_commune=communes.id 
							WHERE voies.nom=?
						  `, voie)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var communes util.Communes
	for rows.Next() {
		var commune util.Commune

		err = rows.Scan(&commune.Id, &commune.Nom, 
						&commune.Cp, &commune.Voies)
		if err != nil {
			log.Fatal(err)
		} 
	
		communes = append(communes, commune)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(communes)		
}

func VoieIdHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	vars := mux.Vars(r)
	voie := vars["id"]	

	rows, err := db.Query(`
						  SELECT communes.id, communes.nom, communes.cp, communes.voies 
							FROM voies 
							INNER JOIN liens ON liens.id_voie=voies.id 
							INNER JOIN communes ON liens.id_commune=communes.id 
							WHERE voies.id=?
						  `, voie)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var communes util.Communes
	for rows.Next() {
		var commune util.Commune

		err = rows.Scan(&commune.Id, &commune.Nom, 
						&commune.Cp, &commune.Voies)
		if err != nil {
			log.Fatal(err)
		} 
	
		communes = append(communes, commune)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(communes)		
}

func VoieLikeHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	vars := mux.Vars(r)
	request := vars["request"]

	// we use Sprintf to properly escape LIKE pattern
	query := fmt.Sprintf("SELECT id, nom, occurences FROM voies WHERE nom LIKE '%%%s%%'", request)
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var voies util.Voies
	for rows.Next() {
		var voie util.Voie

		err = rows.Scan(&voie.Id, &voie.Nom, &voie.Occurences)
		if err != nil {
			log.Fatal(err)
		} 
	
		voies = append(voies, voie)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(voies)		 
}
