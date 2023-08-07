/*
Una interfaz en Golang es como una lista de cosas que le decimos a nuestros objetos que deben saber hacer.

La función principal de una interfaz es ayudarnos a organizar y trabajar con diferentes objetos de una manera más fácil. Nos permite escribir código que funcione con muchos objetos diferentes, siempre que cumplan con las reglas de la interfaz.
*/
package main

import "fmt"

// Definición de la estructura "Person"
type Person struct {
	name string
	age  int
}

// Definición de la estructura "Employee"
type Employee struct {
	id int
}

// Definición de la estructura "FullTimeEmployee" que incorpora "Person" y "Employee"
type FullTimeEmployee struct {
	Person
	Employee
	endDate string
}

// Método para obtener un mensaje de un "FullTimeEmployee"
func (ftEmployee FullTimeEmployee) getMessage() string {
	return "Full Time Employee"
}

// Definición de la estructura "TemporaryEmployee" que incorpora "Person" y "Employee"
type TemporaryEmployee struct {
	Person
	Employee
	taxRate int
}

// Método para obtener un mensaje de un "TemporaryEmployee"
func (tEmployee TemporaryEmployee) getMessage() string {
	return "Temporary Employee"
}

// Definición de la interfaz "PrintInfo"
type PrintInfo interface {
	getMessage() string //solo un metodo que devuelve un string
}

// Función que muestra el mensaje de la interfaz "PrintInfo"
func getMessage(p PrintInfo) {
	fmt.Println(p.getMessage())
}

func main() {
	// Creando una instancia de "FullTimeEmployee"
	ftEmployee := FullTimeEmployee{}

	// Creando una instancia de "TemporaryEmployee"
	tEmployee := TemporaryEmployee{}

	// Llamando a la función "getMessage" para mostrar el mensaje de "TemporaryEmployee"
	getMessage(tEmployee)

	// Llamando a la función "getMessage" para mostrar el mensaje de "FullTimeEmployee"
	getMessage(ftEmployee)
}

/*
En resumen, la interfaz PrintInfo actúa como un contrato que define un método común (getMessage()) que varios structs diferentes (FullTimeEmployee y TemporaryEmployee) deben implementar. Esto nos permite tratar estos structs de manera uniforme y usar la función getMessage() con cualquiera de ellos sin importar sus diferencias internas. La interfaz proporciona una forma de abstraer y unificar el comportamiento de diferentes tipos en Golang.
*/
