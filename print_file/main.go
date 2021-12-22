package main

import (
	"os"
	"fmt"
	"io"
)

func main() {
	path := os.Args[1]

	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	io.Copy(os.Stdout, file)
	// b, err := ioutil.ReadAll(file)
	// fmt.Println(string(b))
}