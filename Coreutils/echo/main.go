package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var n bool
	flag.BoolVar(&n, "n", false, "omit trailing newlines")
	flag.Parse()
	i := 1
	if n {
		i++
		for ; i < len(os.Args); i++ {
			if os.Args[i] != "\n" {
				break
			}
		}
	}
	for ; i < len(os.Args); i++ {
		fmt.Print(os.Args[i])
		if i == len(os.Args)-1 {
			fmt.Println()
		} else {
			fmt.Print(" ")
		}
	}
}
