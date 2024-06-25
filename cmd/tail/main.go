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
	var lines []string
	for in.Scan() {
		lines = append(lines, in.Text())
	}
	if (numberOfLines) > len(lines) {
		log.Fatal("Lines Aren't enough")
	}
	for i := len(lines) - (numberOfLines); i < len(lines); i++ {
		fmt.Println(lines[i])
	}

}
