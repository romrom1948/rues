// Copyright (C) 2016 romrom@tutanota.com
// Use of this source code is governed by the GPLv3
// license that can be found in the LICENSE file.

package util

import (
	"strconv"
)

type Commune struct {
	Id int
	Nom string
	Cp	string
	Voies int
}

// for use with csv.Writer
func (c Commune) Record() []string {
	r := make([]string, 4)
	r[0] = strconv.Itoa(c.Id); r[1] = c.Nom;
	r[2] = c.Cp; r[3] = strconv.Itoa(c.Voies);
	return r 
}

type Voie struct {
	Id int
	Nom string
	Occurences int
}

// for use with csv.Writer
func (v Voie) Record() []string {
	r := make([]string, 3)
	r[0] = strconv.Itoa(v.Id); r[1] = v.Nom; r[2] = strconv.Itoa(v.Occurences);
	return r 
}

type Lien struct {
	Id_voie int
	Id_commune int
	Lat,Long string
}

// for use with csv.Writer
func (l Lien) Record() []string {
	r := make([]string, 4)
	r[0] = strconv.Itoa(l.Id_voie); r[1] = strconv.Itoa(l.Id_commune);
	r[2] = l.Lat; r[3] = l.Long;
	return r 
}
