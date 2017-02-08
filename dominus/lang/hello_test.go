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

	//EE: string to []byte
	s1 := "abcd"
	b1 := []byte(s1)
	fmt.Println(b1) // [97 98 99 100]
	s2 := "中文"
	b2 := []byte(s2)
	fmt.Println(b2) // [228 184 173 230 150 135], unicode，每个中文字符会由三个byte组成
}

func changeName1(p Person) {
	p.firstName = "Bob"
}
func changeName2(p *Person) {
	p.firstName = "Bob"
}
