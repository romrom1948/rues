// Copyright (C) 2016 romrom@tutanota.com
// Use of this source code is governed by the GPLv3
// license that can be found in the LICENSE file.

package main

import (
	"strings"
	"fmt"
	"os"
	"bufio"
	"io"
	"encoding/csv"

	"github.com/romrom1948/rues/util"
)

var helpMessage = []string{
	"Usage : rues_extract <file>",
	"Extract from bano <file> CSV files ready for importation in SQLite.",
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println(strings.Join(helpMessage, "\n"))
		os.Exit(-1)
	}

	err := extract(os.Args[1])			

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

// return both Writer and File so that the caller can manipulate both
func openCSV(path string) (error, *os.File, *csv.Writer) {
	f,err := os.Create(path);	
	if err != nil {
		fmt.Println("impossible to write to" + path)
		return err, nil, nil
	}
	io_w := bufio.NewWriter(f)
	return nil, f, csv.NewWriter(io_w)
}

// create voies.csv, communes.csv and liens.csv
// the structure of the CSV is defined in util/db_format.go
func extract(path string) (e error) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("can not open given input file")
		return err
	}	
	defer func () {
		if err := file.Close(); err != nil {
			fmt.Println("error while closing input file")
		}
	}()
	b := bufio.NewReader(file)
	data := csv.NewReader(b)
	
	e, voies_file, voies_data := openCSV("voies.csv")
	if err != nil {
		fmt.Println("can not voies.csv for writing")
		return err
	}	
	
	e, communes_file, communes_data := openCSV("communes.csv")
	if err != nil {
		fmt.Println("can not communes.csv for writing")
		return err
	}		
	
	e, liens_file, liens_data := openCSV("liens.csv")
	if err != nil {
		fmt.Println("can not communes.csv for writing")
		return err
	}	
	
	// flush everything since the Writer is buffered and report
	// errors on close
	defer func () {
		voies_data.Flush() ; communes_data.Flush() ; liens_data.Flush()
		
		if err := voies_file.Close(); err != nil {
			fmt.Println("error while closing voies.csv")
		}
		
		if err := communes_file.Close(); err != nil {
			fmt.Println("error while closing communes.csv")
		}
		
		if err := liens_file.Close(); err != nil {
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
