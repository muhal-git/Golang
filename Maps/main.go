package main

import "fmt"

func main() {

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}

	colors["green"] = "000000"
	colors["green"] = "123456"

	printMap(colors)

	fmt.Println("colors", colors)

	var colors1 map[string]string
	fmt.Println("colors1", colors1)

	colors2 := make(map[string]string)
	fmt.Println("colors2", colors2)
	colors2["white"] = "#ffffff"
	fmt.Println("colors2", colors2)
	// deleting item
	delete(colors2, "white")
	fmt.Println("colors2", colors2)

}

func printMap(colors map[string]string) {

	for key, value := range colors {
		fmt.Println("Color: ", key, "Hex: ", value)
	}

}
