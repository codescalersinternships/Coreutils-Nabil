package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
func tree(currLevel int, maxLevel int, dirPath string) {
	if currLevel > maxLevel {
		return
	}
	content, err := os.ReadDir(dirPath)
	check(err)

	for _, entry := range content {
		for i := 1; i < currLevel; i++ {
			fmt.Print(" ")
		}
		if entry.IsDir() {
			fmt.Println("|--", entry.Name())
			tree((currLevel + 1), maxLevel, filepath.Join(dirPath, entry.Name()))
		} else {
			fmt.Println("|--", entry.Name())
		}
	}
}

func main() {
	var depthLvl int
	flag.IntVar(&depthLvl, "l", 2, "depth")
	flag.Parse()
	if len(os.Args) <= 1 {
		log.Fatal("No file passed")
		os.Exit(1)
	}
	var dirPath string = os.Args[len(os.Args)-1]

	tree(1, depthLvl, dirPath)
}
