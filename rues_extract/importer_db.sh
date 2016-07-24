#!/bin/sh
# importer les fichiers CSV générés par la commande db dans une base sqlite3
# usage importer_db.sh base.db communes.csv voies.csv liens.csv

if [ "$#" -ne 4 ]; then
	echo "usage: importer_db.sh base.db communes.csv voies.csv liens.csv"
	exit 1
fi

sqlite3 "$1" << EOS

create table communes (id integer, nom text, cp text, voies integer);
create table voies (id integer, nom text, occurences integer);
create table liens (id_voie integer, id_commune integer, lat integer, long integer);
.mode csv
.separator ","
.import "$2" communes
.import "$3" voies
.import "$4" liens
create unique index communes_idx on communes (id);
create unique index communes_nom_idx on communes (nom);
create unique index voies_idx on voies(id);
create unique index voies_nom_idx on voies (nom);

EOS

# on nettoie en cas d'erreur
if [ "$?" -ne 0 ]; then
	rm "$1"
	exit 1
fi
