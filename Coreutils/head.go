package coreutils

import (
	"bufio"
	"fmt"
	"os"
)

func Head(fileOrStd bool, n int, filepath string) {

	in := bufio.NewScanner(os.Stdin)
	if fileOrStd {
		file, err := os.Open(filepath)
		Checkfile(err)
		defer file.Close()
		in = bufio.NewScanner(file)
	}
	var cnt int = 0
	for in.Scan() {
		fmt.Println(in.Text())
		cnt++
		if cnt == (n) {
			break
		}
	}

}
