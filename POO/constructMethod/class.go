package main

import "fmt"

type Employee struct {
	id       int
	name     string
	vacation bool
}

func NewEmployee(id int, name string, vacation bool) *Employee {
	return &Employee{
		id:       id,
		name:     name,
		vacation: vacation,
	}
}

func main() {
	//Empty Object (any value is null, they are initializated with 0 and "")
	e := Employee{}
	fmt.Printf("%v\n", e)

	//Construct (like other languages do)
	e2 := Employee{id: 1, name: "Eduardo", vacation: true}
	fmt.Printf("%v\n", e2)

	//new
	e3 := new(Employee) //return a reference
	fmt.Printf("%v\n", *e3)

	//though a function
	e4 := NewEmployee(1, "Name", true)
	fmt.Printf("%v\n", *e4)
}
