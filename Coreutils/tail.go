package coreutils

import (
	"bufio"
	"fmt"
	"os"
)

func Tail(fileOrStd bool, n int, filepath string) {
	in := bufio.NewScanner(os.Stdin)
	if fileOrStd {
		file, err := os.Open(filepath)
		Checkfile(err)
		defer file.Close()
		in = bufio.NewScanner(file)
	}
	var lines []string
	for in.Scan() {
		lines = append(lines, in.Text())
	}
	if (n) > len(lines) {
		panic("Lines Aren't enough")
	}
	for i := len(lines) - (n); i < len(lines); i++ {
		fmt.Println(lines[i])
	}

}
