# rues

Outils de manipulation de la base de données libre Bano des voies (rues,
avenues, chemins, ...) de France.

* rues_db manipule les fichiers .csv originaux et permet (notamment) de
les passer sous un format importable dans un base SQL
* rues_backend donne accès (en JSON) à la base SQL générée selon une 
API REST-like
* agreger_bano.sh et importer_db.sh permettent de manipuler les fichiers
générés  

# Dépendances

* go
* sqlite3
* [mattn sqlite3 go driver](https://github.com/mattn/go-sqlite3)

# License

GPLv3. Se reporter au fichier LICENSE.
