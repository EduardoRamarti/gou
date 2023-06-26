package main

import "fmt"

func main() {
	var nombre string
	var edad int
	var altura float32

	fmt.Println("Ingresa tu nombre: ")
	fmt.Scanf("%s", &nombre) // %S -> string, &variable -> puntero en el que vamos a guardar lo que recibimos

	fmt.Println(("Ingresa tu edad: "))
	fmt.Scanf("%d", &edad)

	fmt.Println(("Ingresa tu altura: "))
	fmt.Scanf("%f", &altura)

	fmt.Printf("Hola %s con una edad %d y una altura %.2f\n", nombre, edad, altura)
	/*
		La parte %f indica que se imprimirá un valor de punto flotante. El modificador .2 en %f especifica que se mostrarán dos dígitos decimales después del punto decimal.
	*/

}
