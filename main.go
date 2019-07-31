package main

import (
	"path/filepath"

	"github.com/takatoh/mkdocindex/indexmaker"
)


func main() {
	homeDir, _ := filepath.Abs(".")

	maker := indexmaker.New(homeDir)
	maker.Make()
}
