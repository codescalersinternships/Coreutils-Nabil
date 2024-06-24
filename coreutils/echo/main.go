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
	}
	for i, arg := range os.Args {
		if i == 0 || (i == 1 && n) {
			continue
		}
		fmt.Print(arg)
		if i == len(os.Args)-1 {
			if !n {
				fmt.Println()
			}
		} else {
			fmt.Print(" ")
		}
	}
}
