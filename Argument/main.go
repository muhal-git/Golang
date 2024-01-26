package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	f, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Println("Error: ", err)
	}
	io.Copy(os.Stdout, f)
}
