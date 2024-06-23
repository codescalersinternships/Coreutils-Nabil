package coreutils

import (
	"fmt"
	"os"
)

func Echo(n bool) {
	argsWithProg := os.Args
	if n {
		argsWithProg = os.Args[1:]
	}
	fmt.Println(argsWithProg)
}
