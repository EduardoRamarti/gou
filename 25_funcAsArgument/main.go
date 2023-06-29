package main

import "fmt"

// Creando mi tipo de dato
type Operacion func(balance, cantidad int) int //Le estoy indicando a go que cualquier funcion que posea dos parametros de tipo entero y retorne un valor de tipo entero, sera considerada una funcion de tipo operacion

func suma(num1, num2 int) int {
	return num1 + num2
}

func resta(num1, num2 int) int {
	return num1 - num2
}

func ejecutarOperacion(funcion Operacion, balance, cantidad int) {

	fmt.Println("Antes de la operacion")

	resultado := funcion(balance, cantidad)
	fmt.Println("El resultado es:", resultado)

	fmt.Println("Despues de la operacion")

}

func main() {

	ejecutarOperacion(suma, 100, 50)
	ejecutarOperacion(resta, 100, 50)

}
