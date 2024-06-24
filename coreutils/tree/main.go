package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func tree() {
	var dirPath string
	if len(os.Args) == 1 {
		fmt.Println("No file passed")
		os.Exit(1)
	} else {
		dirPath = os.Args[1]
	}
	flag.Parse()

	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			panic(err)
		}
		if info.IsDir() {
			fmt.Printf("dir name: %s\n", path)
		} else {
			fmt.Printf("file name: %s\n", path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
}

func main() {
	tree()
}
