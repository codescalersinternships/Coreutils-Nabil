package coreutils

import (
	"fmt"
	"os"
	"strings"
)

func Env() {
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair[0], ":", pair[1])
	}
}
