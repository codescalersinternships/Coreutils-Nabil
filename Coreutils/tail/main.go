package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func checkFile(e error) {
	if e != nil {
		fmt.Printf("Incorrect file path or file does not exist %e", e)
	}
}

func main() {
	if len(os.Args) == 0 {
		fmt.Println("No file passed")
		os.Exit(1)
	}
	var filepath string = os.Args[len(os.Args)-1]
	n := flag.Int("n", 10, "number of lines")
	flag.Parse()
	file, err := os.Open(filepath)
	checkFile(err)
	defer file.Close()
	in := bufio.NewScanner(file)
	var lines []string
	for in.Scan() {
		lines = append(lines, in.Text())
	}
	if (*n) > len(lines) {
		panic("Lines Aren't enough")
	}
	for i := len(lines) - (*n); i < len(lines); i++ {
		fmt.Println(lines[i])
	}

}
