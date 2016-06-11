// Command base65536-encode encodes its standard input as base65536.
//
// Usage:
//
//     base65536-encode <FILE
//
package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

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
	buf := &bytes.Buffer{}
	first := true
	for _, word := range base65536.Encode(data) {
		if first {
			first = false
		} else {
			buf.WriteString(" ")
		}
		buf.WriteString(word)
		if buf.Len() > 70 {
			buf.WriteString("\n")
			buf.WriteTo(os.Stdout)
			buf.Reset()
			first = true
		}
	}
	buf.WriteTo(os.Stdout)
}
