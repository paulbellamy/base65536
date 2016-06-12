// Command base65536-decode decodes its standard input as base65536. Commas and newlines are ignored.
//
// Usage:
//
//     base65536-decode <FILE
//
package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/paulbellamy/base65536"
)

var prog = filepath.Base(os.Args[0])

func usage() {
	fmt.Fprintf(os.Stderr, "Usage:\n")
	fmt.Fprintf(os.Stderr, "  %s <FILE\n", prog)
	flag.PrintDefaults()
}

func main() {
	log.SetFlags(0)
	log.SetPrefix(prog + ": ")

	flag.Usage = usage
	flag.Parse()

	if flag.NArg() != 0 {
		usage()
		os.Exit(1)
	}

	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	bs, err := base65536.Decode(strings.Fields(string(data)))
	if err != nil {
		log.Fatal(err)
	}
	io.Copy(os.Stdout, bytes.NewReader(bs))
}
