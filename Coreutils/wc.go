package coreutils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Checkfile(e error) {
	if e != nil {
		fmt.Println("Incorrect file path or file does not exist")
		panic(e)
	}
}

func Wc(fileOrStd bool, l bool, w bool, c bool, filepath string) {
	in := bufio.NewScanner(os.Stdin)
	if fileOrStd {
		file, err := os.Open(filepath)
		Checkfile(err)
		defer file.Close()
		in = bufio.NewScanner(file)
	}
	lineCount := 0
	wordCount := 0
	charactersCount := 0
	for in.Scan() {
		lineCount++
		charactersCount += len(in.Text())
		words := strings.Split(in.Text(), " ")
		wordCount += len(words)
	}
	if l {
		fmt.Println(lineCount)
	}
	if w {
		fmt.Println(wordCount)
	}
	if c {
		fmt.Println(charactersCount)
	}

}
