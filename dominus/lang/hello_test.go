package main

import (
	"fmt"
	"testing"
)

var i = 1

type Person struct {
	firstName string
	lastName  string
}

func TestBasic(t *testing.T) {

	//EE: short variable
	i, err := 2, true
	fmt.Println(i, err)

	person := Person{
		firstName: "Alice",
		lastName: "Dow",
	}
	//EE: Passing by value
	changeName1(person)
	fmt.Println(person)
	//EE: Passing by pointer
	changeName2(&person)
	fmt.Println(person)
}

func changeName1(p Person) {
	p.firstName = "Bob"
}
func changeName2(p *Person) {
	p.firstName = "Bob"
}
