package main

import (
	"fmt"
	"reflect"
)

func main() {
	// var course string = "Go Course" //forma 1
	// var course = "Go Course" //forma 2
	course := "Go Course" //forma 3

	longitud := len(course) // devuelve un entero de la cantidad de caracteres dentro de ese string

	fmt.Println(course)
	fmt.Println(longitud)

	primerCaracter := course[0]             // primer caracter
	ultimoCaracter := course[len(course)-1] //ultimo caracter

	fmt.Println(primerCaracter)
	fmt.Println(reflect.TypeOf(primerCaracter)) //uint8 -> representa un char

	/*
		fmt.Printf() se utiliza para formatear e imprimir valores con mayor control. Permite utilizar especificaciones de formato, como %d para enteros, %f para números de punto flotante, %s para cadenas, entre otros.
	*/
	fmt.Printf("%c\n", primerCaracter)
	fmt.Printf("%c\n", ultimoCaracter)
}
