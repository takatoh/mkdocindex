package main

import (
	"fmt"
	"os"
	"path/filepath"
)


func main() {
	var files []string
	var directories []string

	homeDir := "."

	entries, _ := filepath.Glob(homeDir + "/*")

	for _, f := range entries {
		finfo, _ := os.Stat(f)
		if finfo.IsDir() {
			directories = append(directories, f)
		} else {
			files = append(files, f)
		}
	}

	fmt.Println("Directories:")
	for _, d := range directories {
		fmt.Println(d)
	}
	fmt.Println("")
	fmt.Println("Files:")
	for _, f := range files {
		fmt.Println(f)
	}
}
