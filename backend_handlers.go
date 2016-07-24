// Copyright (C) 2016 romrom@tutanota.com
// Use of this source code is governed by the GPLv3
// license that can be found in the LICENSE file.

package rues

import (
    "log"
    "net/http"
	"database/sql"
	"encoding/json"
	
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	
	"github.com/romrom1948/rues/util"	
)

func CommunesHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	rows, err := db.Query("select id, nom, cp, voies from communes")
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

func VoiesHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	rows, err := db.Query("select id, nom, occurences from voies")
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

func CommuneHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	vars := mux.Vars(r)
	commune := vars["commune"]

	rows, err := db.Query(`
						  SELECT voies.id, voies.nom, voies.occurences 
							FROM communes
							INNER JOIN liens ON liens.id_commune=communes.id
							INNER JOIN voies ON liens.id_voie=voies.id
						    WHERE communes.nom=?
						  `, commune)
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


func VoieHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
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
