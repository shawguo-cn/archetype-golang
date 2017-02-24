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

	//short variable
	i, err := 2, true
	fmt.Println(i, err)

	person1 := Person{
		firstName: "Alice",
		lastName: "Dow",
	}

	//In Go assigning one variable to another one creates a copy of the source value and no reference between them is kept.
	person2 := person1; person2.firstName = "shawguo"
	println(person1.firstName)


	//Passing by value
	changeName1(person1)
	fmt.Println(person1)
	//Passing by pointer
	changeName2(&person1)
	fmt.Println(person1)

	fmt.Printf("&person(person address):  %p\n", &person1)

	//string to []byte
	s1 := "abcd"
	b1 := []byte(s1)
	fmt.Println(b1) // [97 98 99 100]
	s2 := "中文"
	b2 := []byte(s2)
	fmt.Println(b2) // [228 184 173 230 150 135], unicode，每个中文字符会由三个byte组成
}

func TestPointer(t *testing.T) {
	var a int   // a variable of type 'int'
	var b int   // another 'int'
	var c *int  // a variable of type 'pointer to int'

	a = 42      // assigning value '42' to 'a'
	b = a       // copying the value from 'a' to 'b'
	c = &a      // assigning the address of 'a' to 'c'
	fmt.Printf("Values:\n;a = %v\n b = %v\n c = %v\n*c = %v\n", a, b, c, *c)


	a = 21      // assigning a new value of '21' to 'a'
	fmt.Printf("Values:\n a = %v\n b = %v\n c = %v\n*c = %v\n", a, b, c, *c)
	fmt.Printf("Types:\n a:  %T\n b:  %T\n c: %T\n*c:  %T\n", a, b, c, *c)
}

func changeName1(p Person) {
	p.firstName = "Bob"
}
func changeName2(p *Person) {
	p.firstName = "Bob"
}
