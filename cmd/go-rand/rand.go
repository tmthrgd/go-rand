// Copyright 2018 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License license that can be found in
// the LICENSE file.

package main

import (
	"flag"
	"io"
	"log"
	"os"

	"github.com/tmthrgd/go-rand"
)

func main() {
	var n uint
	flag.UintVar(&n, "c", 0, "the number of bytes to read")

	flag.Parse()

	var limit bool
	flag.Visit(func(f *flag.Flag) {
		if f.Name == "c" {
			limit = true
		}
	})

	var err error
	if limit {
		_, err = io.CopyN(os.Stdout, rand.Reader, int64(n))
	} else {
		_, err = io.Copy(os.Stdout, rand.Reader)
	}

	if err != nil {
		log.Fatal(err)
	}
}
