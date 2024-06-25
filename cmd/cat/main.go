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
	var printLinesNumber bool
	flag.BoolVar(&printLinesNumber, "n", false, "number of lines")
	flag.Parse()
	if len(os.Args) < 2 {
		log.Fatal("No file passed")
	}
	var filepath string = os.Args[len(os.Args)-1]
	file, err := os.Open(filepath)
	checkFileError(err)
	defer file.Close()
	in := bufio.NewScanner(file)
	var cnt int = 0
	for in.Scan() {
		if printLinesNumber {
			fmt.Printf("%d ", cnt)
		}
		fmt.Println(in.Text())
		cnt++
	}
}
