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
		os.Exit(1)
	}
}

func main() {
	var n bool
	flag.BoolVar(&n, "n", false, "number of lines")
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
		if n {
			fmt.Printf("%d ", cnt)
		}
		fmt.Println(in.Text())
		cnt++
	}
}
