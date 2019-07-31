package main

import (
//	"fmt"
//	"os"
	"path/filepath"

	"github.com/takatoh/mkdocindex/indexmaker"
)


func main() {
//	var files []string
//	var directories []string

	homeDir, _ := filepath.Abs(".")

	maker := indexmaker.New(homeDir)
	maker.Make()
}
