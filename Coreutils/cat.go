package coreutils

import (
	"bufio"
	"fmt"
	"os"
)

func Cat(filepath string) {
	file, err := os.Open(filepath)
	Checkfile(err)
	defer file.Close()
	in := bufio.NewScanner(file)
	var cnt int = 0
	for in.Scan() {
		fmt.Println(cnt, in.Text())
		cnt++
	}
}
