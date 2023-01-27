package main

import (
	"fmt"
	"os"
	"strings"
)

var progname = basename(os.Args[0], "")

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Usage:", progname, "name [suffix]")
		os.Exit(1)
	}
	suffix := ""
	if len(os.Args) > 2 {
		suffix = os.Args[2]
	}
	fmt.Println(basename(os.Args[1], suffix))
}

func basename(name, suffix string) string {
	for name[len(name)-1] == '/' || name[len(name)-1] == '\\' {
		if len(name) == 1 {
			return name
		}
		name = name[:len(name)-1]
	}
	name = name[strings.LastIndexAny(name, "/\\")+1:]
	if name != suffix && strings.HasSuffix(name, suffix) {
		name = name[:strings.LastIndex(name, suffix)]
	}
	return name
}
