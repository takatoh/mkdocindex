package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/takatoh/mkdocindex/indexmaker"
)


const (
	progVersion = "v0.2.2"
)


func main() {
	opt_version := flag.Bool("v", false, "Show version.")
	flag.Parse()

	if *opt_version {
		fmt.Println(progVersion)
		os.Exit(0)
	}

	var dir string
	if len(flag.Args()) > 0 {
		dir = flag.Args()[0]
		fmt.Println(dir)
	} else {
		dir = "."
	}
	homeDir, _ := filepath.Abs(dir)

	maker := indexmaker.New(homeDir)
	maker.Make()
}
