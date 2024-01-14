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

	// instead we use pointers to update something
	/*
		Turn value into address with &value
		Turn address into value with *address
	*/
	jimPointer := &jim
	jimPointer.updateName2("jimmy")
	jim.print()

	/*
		instead of saying jimPointer := &jim and then jimPointer.updateName2("jimmy")
		wen can simply say jim.updateName2("jimmy"). This will also update the name
	*/

	/*
		In GO everything is passed by value.
		When we use Value Types we need to use pointers to update variables inside the functions,
		but when we use Reference Types we dont need to worry about pointers.

		slices, maps, channels, pointers, functions are Reference types.
	*/

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
