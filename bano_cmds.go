// Copyright (C) 2016 romrom@tutanota.com
// Use of this source code is governed by the GPLv3
// license that can be found in the LICENSE file.

package rues

import (
	"fmt"
	"io"
	"bufio"
	"os"
	"sort"
	"strings"
	"encoding/csv"
	
	"github.com/romrom1948/rues/util"
)

// print street names ranked by frequency	
func Ranking(r io.Reader) (e error) {
	data := csv.NewReader(r)

	// m1: (street name -> list of towns it is in)
	// m2: (street name -> number of towns it is in)
	// m3: m2 reversed (# of occurences -> names)
	// we build the three maps in parallel to avoid looping 3x on the data
	var m1 = make(map[string][]string)
	var m2 = make(map[string]int)
	var m3 = make(map[int][]string)
	
	for {
		record, err := data.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		if ! util.IsIn(record[4], m1[record[2]]) {
			m1[record[2]] = append(m1[record[2]], record[4])

			if m2[record[2]] == 0 { // first occurence
				m2[record[2]]++
				m3[m2[record[2]]] = append(m3[m2[record[2]]], record[2])
			} else { // already seen once
				i, e := util.IsAt(record[2], m3[m2[record[2]]])
				if e != nil { // by construction, record[2] should be somewhere
							  // in the slice ; something is wrong if not
					fmt.Println("internal error");
					return e
				}				
				// move the name according to the new # of occurences
				m3[m2[record[2]]] = append(m3[m2[record[2]]][:i], 
										   m3[m2[record[2]]][i+1:]...)
				m3[m2[record[2]]+1] = append(m3[m2[record[2]]+1], record[2])

				m2[record[2]]++
			}
		}
	}
	
	// print in an ordered way
	var keys []int
	for k := range m3 {
		if len(m3[k]) != 0 {
			keys = append(keys, k)
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))
	for i, k := range keys {
		fmt.Printf("%d [%d voie(s)] - %s (%d occurences)\n", 
					i+1, len(m3[k]), strings.Join(m3[k], ", "), k)
	}

	return nil
}

// create voies.csv, communes.csv and liens.csv
// the structure of the CSV is defined in util/db_format.go
func Extract(r io.Reader) (e error) {
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
				vid++ // update now, create the link after
			} else {
				v.Occurences++
			}
			
			c := cmap[record[4]]
			if c == nil { // create if necessary
				cmap[record[4]] = &util.Commune{cid, record[4], record[3], 1}
				cidx[record[4]] = cid
				cid++
			} else {
				c.Voies++ // update now, create the link after
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
