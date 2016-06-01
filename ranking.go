package rues

import (
	"fmt"
	"sort"
	"io"
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
