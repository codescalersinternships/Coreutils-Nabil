package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	var n bool
	flag.BoolVar(&n, "n", false, "omit trailing newlines")
	flag.Parse()
	startIndex := 1

	if n {
		startIndex = 2
	}

	args := os.Args[startIndex:]
	output := strings.Join(args, " ")

	if n {
		fmt.Print(output)
	} else {
		fmt.Println(output)
	}
}
