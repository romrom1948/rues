package rues

import (
	"fmt"
	"io"
	"bufio"
	"os"
	"encoding/csv"
	
	"github.com/romrom1948/rues/util"
)

// create voies.csv, communes.csv and liens.csv
// the structure of the CSV is defined above
func DB(r io.Reader) (e error) {
	data := csv.NewReader(r)
	
	fvoies,err := os.Create("voies.csv");	
	if err != nil {
		fmt.Println("impossible to write to voies.csv")
		return err
	}
	wvoies := bufio.NewWriter(fvoies)
	voies_data := csv.NewWriter(wvoies);
	
	fcommunes,err := os.Create("communes.csv");	
	if err != nil {
		fmt.Println("impossible to write to communes.csv")
		return err
	}
	wcommunes := bufio.NewWriter(fcommunes)
	communes_data := csv.NewWriter(wcommunes);
	
	fliens,err := os.Create("liens.csv");	
	if err != nil {
		fmt.Println("impossible to write to liens.csv")
		return err
	}
	wliens := bufio.NewWriter(fliens)
	liens_data := csv.NewWriter(wliens);
	
	// we shouldn't forget to flush everything since the Writer is
	// buffered
	defer func () {
		voies_data.Flush() ; communes_data.Flush() ; liens_data.Flush()
		
		if err := fvoies.Close(); err != nil {
			fmt.Println("error while closing voies.csv")
		}
		
		if err := fcommunes.Close(); err != nil {
			fmt.Println("error while closing communes.csv")
		}
		
		if err := fliens.Close(); err != nil {
			fmt.Println("error while closing liens.csv")
		}	
	}()
	
	var vmap = make(map[string]*util.Voie)
	var cmap = make(map[string]*util.Commune)
	var liens []*util.Lien	
	
	var vidx = make(map[string]int);
	var cidx = make(map[string]int);
	
	var seen = make(map[string][]string)
	var vid int = 1
	var cid int = 1

	for {
		record, err := data.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}
											
		if ! util.IsIn(record[4], seen[record[2]]) { // new link
			seen[record[2]] = append(seen[record[2]], record[4])

			v := vmap[record[2]]
			if v == nil { // create if necessary
				vmap[record[2]] = &util.Voie{vid, record[2], 1}
				vidx[record[2]] = vid
				vid++ // update even if the link is created after
			} else {
				v.Occurences++
			}
			
			c := cmap[record[4]]
			if c == nil { // create if necessary
				cmap[record[4]] = &util.Commune{cid, record[4], record[3], 1}
				cidx[record[4]] = cid
				cid++
			} else {
				c.Voies++ // update even if the link is created after
			}
			
			liens = append(liens, &util.Lien{vidx[record[2]], cidx[record[4]], record[6], record[7]})			
		}
	}
	
	for _, v := range vmap {
		err := voies_data.Write(v.Record())
		if err != nil {
			fmt.Println("error while writing to voies.scv")
			return err
		}
	}
	
	for _, c := range cmap {
		err := communes_data.Write(c.Record())
		if err != nil {
			fmt.Println("error while writing to communes.scv")
			return err
		}
	}		
	
	for _, l := range liens {
		err := liens_data.Write(l.Record())
		if err != nil {
			fmt.Println("error while writing to liens.scv")
			return err
		}			
	}
	
	return nil
}


