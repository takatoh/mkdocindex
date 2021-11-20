package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/takatoh/mkdocindex/indexmaker"
)

const (
	progVersion = "v0.6.0"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr,
			`Usage:
  %s [options] [dir]

Options:
`, os.Args[0])
		flag.PrintDefaults()
	}
	opt_distributed := flag.Bool("d", false, "Generate distributed HTML.")
	opt_version := flag.Bool("v", false, "Show version.")
	flag.Parse()

	if *opt_version {
		fmt.Println(progVersion)
		os.Exit(0)
	}

	var dir string
	if len(flag.Args()) > 0 {
		dir = flag.Args()[0]
	} else {
		dir = "."
	}
	root, _ := filepath.Abs(dir)

	maker := indexmaker.New(root)
	if *opt_distributed {
		maker.Make()
	} else {
		maker.MakeMonolithic()
	}
}
