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

	"github.com/romrom1948/rues"
)

var helpMessage = []string{
	"Usage : rues_db <cmd> <file>",
	"Process bano file <file> according to <cmd>.",
	"",
	"<cmd> must be one of:",
	"	help 		print this help message",
	"	extract		create voies.scv, communes.scv and liens.scv",			
	"	ranking		print street names ranked by frequency",
	"",
}

func abort(msg string) {
	fmt.Println(msg)
	os.Exit(-1)
}

func run(cmd func(r io.Reader) (e error), path string) (e error) {
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

	return cmd(b)
}

func main() {
	var err error

	if len(os.Args) == 1 {
		fmt.Println(strings.Join(helpMessage, "\n"))
		abort("no command given")
	}

	if os.Args[1] == "help" {
		fmt.Println(strings.Join(helpMessage, "\n"))
		os.Exit(0)
	}

	if len(os.Args) == 2 {
		abort("no file given")
	}

	switch os.Args[1] {
		case "db":
			err = run(rues.Extract, os.Args[2])			
		case "ranking":
			err = run(rues.Ranking, os.Args[2])
		default:
			abort("invalid command")
	}

	if err != nil {
		fmt.Println(err)
	}
}
