package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func checkFile(e error) {
	if e != nil {
		fmt.Printf("Incorrect file path or file does not exist %e", e)
	}
}

func main() {
	l := flag.Bool("l", false, "Print Number of lines")
	w := flag.Bool("w", false, "Print Number of words")
	c := flag.Bool("c", false, "Print Number of characters")
	flag.Parse()
	if len(os.Args) == 0 {
		fmt.Println("No file passed")
		os.Exit(1)
	}
	var filepath string = os.Args[len(os.Args)-1]
	file, err := os.Open(filepath)
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
	if *l {
		fmt.Println(lineCount)
	}
	if *w {
		fmt.Println(wordCount)
	}
	if *c {
		fmt.Println(charactersCount)
	}

}
