package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func checkFile(e error) {
	if e != nil {
		log.Fatal("Incorrect file path or file does not exist ", e)
	}
}

func main() {
	var l, w, c bool

	flag.BoolVar(&l, "l", false, "Print Number of lines")
	flag.BoolVar(&w, "w", false, "Print Number of words")
	flag.BoolVar(&c, "c", false, "Print Number of characters")
	flag.Parse()
	if len(os.Args) == 0 {
		log.Fatal("No file passed")
	}
	var filePath string = os.Args[len(os.Args)-1]
	file, err := os.Open(filePath)
	checkFile(err)
	defer file.Close()
	in := bufio.NewScanner(file)
	lineCount := 0
	wordCount := 0
	charactersCount := 0
	for in.Scan() {
		lineCount++
		charactersCount += len(in.Text())
		words := strings.Split(in.Text(), " ")
		wordCount += len(words)
	}
	if l {
		fmt.Println(lineCount)
	}
	if w {
		fmt.Println(wordCount)
	}
	if c {
		fmt.Println(charactersCount)
	}

}
