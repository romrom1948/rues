// Copyright (C) 2016 romrom@tutanota.com
// Use of this source code is governed by the GPLv3
// license that can be found in the LICENSE file.

package main

import (
    "net/http"
	"database/sql"
	"encoding/json"
	
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	
	"github.com/romrom1948/rues/util"	
)

func VoiesHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) (error){
	util.JsonHeader(w)	

	rows, err := db.Query(`SELECT id, nom, occurences FROM voies`)
	if err != nil {
		return err
	}
	defer rows.Close()

	var voies util.Voies
	for rows.Next() {
		var voie util.Voie
		if rows.Scan(&voie.Id, &voie.Nom, &voie.Occurences) != nil {
			return err
		}
		voies = append(voies, voie)
	}
	if rows.Err() != nil {
		return err
	}

    w.WriteHeader(http.StatusOK)	
	json.NewEncoder(w).Encode(voies)
	
	return nil
}

func VoieNameHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) (error){
	vars := mux.Vars(r)
	voie := vars["voie"]
		
	util.JsonHeader(w)	

	rows, err := db.Query(`
						  SELECT communes.id, communes.nom, communes.cp, communes.voies 
							FROM voies 
							INNER JOIN liens ON liens.id_voie=voies.id 
							INNER JOIN communes ON liens.id_commune=communes.id 
							WHERE voies.nom=?
						  `, voie)
	if err != nil {
		return err
	}
	defer rows.Close()

	var communes util.Communes
	for rows.Next() {
		var commune util.Commune
		if rows.Scan(&commune.Id, &commune.Nom, 
					 &commune.Cp, &commune.Voies) != nil {
			return err
		}
		communes = append(communes, commune)
	}
	if rows.Err() != nil {
		return err
	}

    w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(communes)
	
	return nil	
}

func VoieIdHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) (error){
	vars := mux.Vars(r)
	voie := vars["id"]
	
	util.JsonHeader(w)	

	rows, err := db.Query(`
						  SELECT communes.id, communes.nom, communes.cp, communes.voies 
							FROM voies 
							INNER JOIN liens ON liens.id_voie=voies.id 
							INNER JOIN communes ON liens.id_commune=communes.id 
							WHERE voies.id=?
						  `, voie)
	if err != nil {
		return err
	}
	defer rows.Close()

	var communes util.Communes
	for rows.Next() {
		var commune util.Commune

		if rows.Scan(&commune.Id, &commune.Nom, 
					 &commune.Cp, &commune.Voies) != nil {
			return err
		}
		communes = append(communes, commune)
	}
	if rows.Err() != nil {
		return err
	}

    w.WriteHeader(http.StatusOK)	
	json.NewEncoder(w).Encode(communes)
	
	return nil
}

func VoieLikeHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) (error){
	vars := mux.Vars(r)
	request := vars["request"]
	
	util.JsonHeader(w)		

	rows, err := db.Query(`SELECT id, nom, occurences FROM voies WHERE nom LIKE ?`, 
						  string('%') + request + string('%')) // ugly
	if err != nil {
		return err
	}
	defer rows.Close()

	var voies util.Voies
	for rows.Next() {
		var voie util.Voie
		if rows.Scan(&voie.Id, &voie.Nom, &voie.Occurences) != nil {
			return err
		}
		voies = append(voies, voie)
	}
	if rows.Err() != nil {
		return err
	}

    w.WriteHeader(http.StatusOK)	
	json.NewEncoder(w).Encode(voies)	
	
	return nil
}
