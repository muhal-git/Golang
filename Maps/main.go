package main

import "fmt"

func main() {

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}

	fmt.Println("colors", colors)

	var colors1 map[string]string
	fmt.Println("colors1", colors1)

	colors2 := make(map[string]string, 5)
	fmt.Println("colors2", colors2)
	colors2["white"] = "#ffffff"
	fmt.Println("colors2", colors2)
	delete(colors2, "white")
	fmt.Println("colors2", colors2)

}
