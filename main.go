package main

import (
	"flag"
	"fmt"
)

var filename string

func Flags() {
	flag.StringVar(&filename, "filename", "", "filename of the ROM to load")
	flag.Parse()
}

func main() {
	Flags()
	fmt.Println(filename)
}
