package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func checkFileError(e error) {
	if e != nil {
		log.Fatal("Incorrect file path or file does not exist ", e)
	}
}
func main() {
	var numberOfLines int
	flag.IntVar(&numberOfLines, "n", 10, "number of lines")
	flag.Parse()
	if len(os.Args) <= 1 {
		log.Fatal("No file passed")
		os.Exit(1)
	}
	var filepath string = os.Args[len(os.Args)-1]
	file, err := os.Open(filepath)
	checkFileError(err)
	defer file.Close()
	in := bufio.NewScanner(file)
	var cnt int = 0
	for in.Scan() {
		fmt.Println(in.Text())
		cnt++
		if cnt == (numberOfLines) {
			break
		}
	}
	if cnt != numberOfLines {
		log.Fatal("Lines weren't enough")
	}

}
