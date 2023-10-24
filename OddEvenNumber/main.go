package main

import "fmt"

func main() {

	numbers := []int{}

	for i := 0; i < 10; i++ {
		numbers = append(numbers, i)
	}

	for _, number := range numbers {

		if number%2 == 0 {
			fmt.Println(number, "is even.")
		} else {
			fmt.Println(number, "is odd.")
		}
	}
}
