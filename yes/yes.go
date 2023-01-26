package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var s string
	if len(os.Args) == 1 {
		s = "y"
	} else {
		s = strings.Join(os.Args, " ")
	}
	for {
		fmt.Println(s)
	}
}
