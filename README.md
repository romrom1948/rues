# rues

Outils de manipulation de la base de données libre Bano des voies (rues,
avenues, chemins, ...) de France.

* rues_db manipule les fichiers .csv originaux et permet (notamment) de
les convertir dans un format importable dans une base SQL
* agreger_bano.sh et importer_db.sh permettent de manipuler les fichiers
générés (compilation et importation dans une base)
* rues_backend est un serveur de backend JSON selon une API rest-like à
partir de la base données crée
* rues_frontend est le frontend web correspondant au backend

# Dépendances

* go
* sqlite3
* [mattn sqlite3 go driver](https://github.com/mattn/go-sqlite3)

# License

GPLv3. Se reporter au fichier LICENSE.
