package main

import (
//	"fmt"
//	"os"
//	"path/filepath"

	"github.com/takatoh/mkdocindex/indexmaker"
)


func main() {
//	var files []string
//	var directories []string

	homeDir := "."

	maker := indexmaker.New(homeDir)
	maker.Make()
}
