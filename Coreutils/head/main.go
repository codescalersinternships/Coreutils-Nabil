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
	var n int
	flag.IntVar(&n, "n", 10, "number of lines")
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
	var cnt int = 0
	for in.Scan() {
		fmt.Println(in.Text())
		cnt++
		if cnt == (n) {
			break
		}
	}
	if cnt != n {
		fmt.Println("Lines weren't enough")
	}

}
