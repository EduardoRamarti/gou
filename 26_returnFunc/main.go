package main

import "fmt"

// creo mi tipo de dato operacion
type Operacion func(num1, num2 int) int

// Creo mi funcion  que recibe un string y retorna un tipo Operacion (funcion)
func crearOperacion(tipo string) Operacion {
	if tipo == "suma" { //si el string es igual a "suma" entonces regresa la sigueinte funcion
		return func(n1, n2 int) int { return n1 + n2 }
	} else if tipo == "resta" {
		return func(num1, num2 int) int { return num1 - num2 }
	} else {
		return func(num1, num2 int) int { return num1 * num2 }
	}
}

func main() {

	suma := crearOperacion("suma")

	resultado := suma(40, 50)

	fmt.Println("El resultado es:", resultado)

}
