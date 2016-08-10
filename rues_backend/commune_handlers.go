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

func CommunesHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) (error){
	util.JsonHeader(w)	

	rows, err := db.Query(`SELECT id, nom, cp, voies FROM communes`)
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

func CommuneNameHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) (error){
	util.JsonHeader(w)	
	
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

func CommuneIdHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) (error){
	util.JsonHeader(w)	
	
	vars := mux.Vars(r)
	commune := vars["id"]

	rows, err := db.Query(`
						  SELECT voies.id, voies.nom, voies.occurences 
							FROM communes
							INNER JOIN liens ON liens.id_commune=communes.id
							INNER JOIN voies ON liens.id_voie=voies.id
						    WHERE communes.id=?
						  `, commune)
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

func CommuneLikeHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) (error){	
	vars := mux.Vars(r)
	request := vars["request"]
	
	util.JsonHeader(w)	
	
	rows, err := db.Query(`SELECT id, nom, cp, voies FROM communes WHERE nom LIKE ?`, 
						  string('%') + request + string('%')) // ugly
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
