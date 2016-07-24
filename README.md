# rues

Outils de manipulation de la base de données libre Bano des voies (rues,
avenues, chemins, ...) de France.

* agreger_bano.sh agrège en un seul fichier les fichiers par département
* rues_extract extrait les informations de la base données CSV de départ
dans un format importable sous sqlite
* importer_db.sh importe ce qui a été extrait dans une base sqlite
* rues_backend est un serveur de backend JSON qui vient se brancher sur
la base sqlite extraite
* enfin, rues_frontend est le frontend web correspondant au backend

# Dépendances

* go
* sqlite3
* [mattn sqlite3 go driver](https://github.com/mattn/go-sqlite3)

# License

GPLv3. Se reporter au fichier LICENSE.
