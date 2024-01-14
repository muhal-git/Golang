package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {

	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// PROBLEM!
/*
func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

func printGreeting(sb spanishBot) {
	fmt.Println(sb.getGreeting())
}
*/

/*
we can remove eb in receiver part (eb englishBot),
we dont use it.
*/
func (englishBot) getGreeting() string {

	return "Hello!"
}

func (sb spanishBot) getGreeting() string {

	return "Hola!"
}
