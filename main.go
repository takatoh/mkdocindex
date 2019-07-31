package main

import (
	"fmt"
//	"os"
	"path/filepath"
)


func main() {
	homeDir := "."

	entries, _ := filepath.Glob(homeDir + "/*")

	for _, f := range entries {
		fmt.Println(f)
	}
}
