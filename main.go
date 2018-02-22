package main

import (
	"flag"
	"fmt"
	"io/ioutil"
)

var filename string

func Flags() {
	flag.StringVar(&filename, "filename", "", "filename of the ROM to load")
	flag.Parse()
}

func main() {
	Flags()
	ROM, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Errorf("%v", err)
	}
	fmt.Println(ROM)
}
