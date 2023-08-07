package main

import "fmt"

/*
Go no permite la herencia, go utiliza la composicion.
la composicion, a diferencia de la herencia, no es una clase hija de… sino que contiene los metodos de las clases indicadas.
*/

type Person struct {
	name string
	age  int
}

type Employee struct {
	id int
}

type FullTimeEmployee struct {
	Person
	Employee
}

func GetMessage(p Person) {
	fmt.Printf("%s with age %d\n", p.name, p.age)
}

func main() {
	ftEmployee := FullTimeEmployee{}
	ftEmployee.name = "Name"
	ftEmployee.age = 2
	ftEmployee.id = 5
	fmt.Printf("%v\n", ftEmployee)

	//GetMessage(ftEmployee)
}
