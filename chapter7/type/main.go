package main

import (
	"fmt"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func (p Person) String() string {
	return fmt.Sprintf("%s %s, age is %d", p.FirstName, p.LastName, p.Age)
}
func main() {
	p := Person{
		FirstName: "John",
		LastName:  "Thomason",
		Age:       29,
	}
	output := p.String()
	fmt.Println(output)
	
}