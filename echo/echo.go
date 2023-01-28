package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	flag.Usage = func() {
		name := os.Args[0][strings.LastIndexAny(os.Args[0], "/\\")+1:]
		fmt.Fprintf(os.Stderr, "Usage: %s [-n] text\n", name)
	}
	n := flag.Bool("n", false, "suppress newline")
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), " "))
	if !*n {
		fmt.Println()
	}
}
