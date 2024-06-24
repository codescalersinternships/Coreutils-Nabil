package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	for _, e := range os.Environ() {
		idx := strings.Index(e, "=")
		key := e[:idx]
		value := e[idx+1:]
		fmt.Println(key, ":", value)
	}
}
