package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	var omitNewline bool
	flag.BoolVar(&omitNewline, "n", false, "omit trailing newlines")
	flag.Parse()
	startIndex := 1

	if omitNewline {
		startIndex = 2
	}

	args := os.Args[startIndex:]
	output := strings.Join(args, " ")

	if omitNewline {
		fmt.Print(output)
	} else {
		fmt.Println(output)
	}
}
