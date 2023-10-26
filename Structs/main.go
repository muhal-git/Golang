package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	age       int
	isAlive   bool
	test      uint64
	contactInfo
}

func main() {
	// ways of creating a new struct
	alex := person{"Alex", "Anderson", 25, false, 5, contactInfo{}}
	fmt.Println(alex)
	// better way
	sarah := person{
		firstName: "Sarah",
		lastName:  "Musk"}
	fmt.Println(sarah)
	// update
	sarah.lastName = "Muska"
	fmt.Println(sarah)

	// third way
	var person1 person
	fmt.Println(person1)

	fmt.Printf("%+v", person1)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		age:       20,
		isAlive:   true,
		test:      125,
		contactInfo: contactInfo{
			email:   "jim@mail",
			zipCode: 12345,
		},
	}

	fmt.Printf("\n%+v", jim)

	jim.print()
	// following will not update jim
	jim.updateName1("jimmy")
	jim.print()
	// instead we use pointers
	jimPointer := &jim
	jimPointer.updateName2("jimmy")
	jim.print()

}

func (p person) print() {

	fmt.Printf("\n%+v\n", p)
}

func (p person) updateName1(newFirstName string) {

	p.firstName = newFirstName
}

func (pointerToPerson *person) updateName2(newFirstName string) {

	(*pointerToPerson).firstName = newFirstName
}
