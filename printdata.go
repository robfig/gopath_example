package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	mustPrintData("foo-data/file.txt")
	mustPrintData("bar-data/file.txt")
}

func mustPrintData(filename string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("error reading %v: %v\n", filename, err)
		os.Exit(1)
	}
	fmt.Println(string(data))
}
