package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/takatoh/mkdocindex/indexmaker"
)


const (
	progVersion = "v0.1.0"
)


func main() {
	opt_version := flag.Bool("v", false, "Show version.")
	flag.Parse()

	if *opt_version {
		fmt.Println(progVersion)
		os.Exit(0)
	}

	homeDir, _ := filepath.Abs(".")

	maker := indexmaker.New(homeDir)
	maker.Make()
}
