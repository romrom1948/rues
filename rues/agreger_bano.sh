#!/bin/sh
# agreger_bano.sh <fichiers>
# produit (stdout) un CSV agrégé de tous les fichiers bano passés en argument
# un id_departement est ajouté en fonction des codes postaux des communes
# exemple: 
#		find ~/bano-data/bano-*.csv -type f -exec agreger_bano.sh {} + 

for FILE in "$@"
do
	DPT=`awk -F',' 'BEGIN { OFS="," } {print $4}' $FILE | head -n 1 | cut -c 1-2`
	# traitement spécial pour les DOM
	if [ "$DPT" -eq "97" ]; then 
		DPT=`awk -F',' 'BEGIN { OFS="," } {print $4}' $FILE | head -n 1 | cut -c 1-3`
	fi	

	sed "s/$/,$DPT/" $FILE
done
