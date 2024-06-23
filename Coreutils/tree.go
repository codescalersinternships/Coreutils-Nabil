package coreutils

import (
	"fmt"
	"os"
	"path/filepath"
)

func Tree(filePath string) {

	err := filepath.Walk(*&filePath, func(path string, info os.FileInfo, err error) error {
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
