package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	links := []string{
		"http://google.com.tr",
		"http://amazon.com",
		"http://metu.edu.tr",
		"http://trt.gov.tr",
	}

	c := make(chan string)

	for _, link := range links {
		// use go keyword before function calls
		go checkLink(link, c)
	}
	/*
		for {
			go checkLink(<-c, c)
		}
	*/
	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	fmt.Println("Checking link... ", link)
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " might be down!")
		fmt.Println("\t Error: ", err)
		c <- link
		return
	}

	fmt.Println(link, " is up!")
	c <- link
}
