package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func checkFile(e error) {
	if e != nil {
		log.Fatal("Incorrect file path or file does not exist ", e)
	}
}

func main() {
	var n int
	flag.IntVar(&n, "n", 10, "number of lines")
	flag.Parse()
	if len(os.Args) <= 1 {
		log.Fatal("No file passed")
		os.Exit(1)
	}
	var filepath string = os.Args[len(os.Args)-1]
	file, err := os.Open(filepath)
	checkFile(err)
	defer file.Close()
	in := bufio.NewScanner(file)
	var lines []string
	for in.Scan() {
		lines = append(lines, in.Text())
	}
	if (n) > len(lines) {
		log.Fatal("Lines Aren't enough")
	}
	for i := len(lines) - (n); i < len(lines); i++ {
		fmt.Println(lines[i])
	}

}
