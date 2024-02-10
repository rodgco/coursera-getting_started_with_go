package main

import "fmt"

type Person struct {
	name string
	age int
	address string
	phone string
}

func main () {
	var person1 Person
	person1.name = "John"
	person1.age = 25
	person1.address = "New York"
	fmt.Println("Age:", person1.age)
	fmt.Println("Address:", person1.address)
	fmt.Println("Phone:", person1.phone)
	person1.phone = "1234567890"

	fmt.Println("Name:", person1.name)
}
