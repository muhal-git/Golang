package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {

	resp, err := http.Get("http://universities.hipolabs.com/search?country=Turkey")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	/*
		bs := make([]byte, 99999)
		resp.Body.Read(bs)

		fmt.Printf("bs: %v\n", string(bs))
	*/

	lw := logWriter{}

	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Written ", len(bs), " bytes.")

	return len(bs), nil
}
