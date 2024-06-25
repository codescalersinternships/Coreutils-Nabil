package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func checkFileError(e error) {
	if e != nil {
		log.Fatal("Incorrect file path or file does not exist ", e)
	}
}

func main() {
	var printLineCount, printWordCount, printCharactersCount bool

	flag.BoolVar(&printLineCount, "l", false, "Print Number of lines")
	flag.BoolVar(&printWordCount, "w", false, "Print Number of words")
	flag.BoolVar(&printCharactersCount, "c", false, "Print Number of characters")
	flag.Parse()
	if len(os.Args) == 0 {
		log.Fatal("No file passed")
	}
	var filePath string = os.Args[len(os.Args)-1]
	file, err := os.Open(filePath)
	checkFileError(err)
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
	if printLineCount {
		fmt.Println(lineCount)
	}
	if printWordCount {
		fmt.Println(wordCount)
	}
	if printCharactersCount {
		fmt.Println(charactersCount)
	}

}
