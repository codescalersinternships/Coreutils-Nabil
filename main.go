package main

import (
	coreutils "Coreutils/Coreutils"
	"flag"
	"fmt"
)

func True() int {
	return 0
}
func False() int {
	return 1
}
func main() {
	operation := flag.String("operation", "none", "operation name")
	filepath := flag.String("directory", "input.txt", "operation name")
	fileOrStd := flag.Bool("isFile", true, "True readfrom file")
	n := flag.Int("n", 10, "number of lines")
	o := flag.Bool("o", false, "Omit trainling newline ")
	l := flag.Bool("l", false, "Print Number of lines")
	w := flag.Bool("w", false, "Print Number of words")
	c := flag.Bool("c", false, "Print Number of characters")
	flag.Parse()
	switch *operation {
	case "head":
		coreutils.Head(*fileOrStd, *n, *filepath)
	case "tail":
		coreutils.Tail(*fileOrStd, *n, *filepath)
	case "cat":
		coreutils.Cat(*filepath)
	case "echo":
		coreutils.Echo(*o)
	case "tree":
		coreutils.Tree(*filepath)
	case "wc":
		coreutils.Wc(*fileOrStd, *l, *w, *c, *filepath)
	case "env":
		coreutils.Env()
	case "true":
		True()
	case "false":
		False()
	default:
		fmt.Println("Check you spelling")
	}
}
