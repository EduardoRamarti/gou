package main

import "fmt"

// se debe indicar el tipo de dato del parametro a introducir
func saluda(username string) {
	fmt.Println("Hola", username)
}

// Retornar valores de una funcion
// tambien se indica el tipo de dato a retornar
func suma(numero1 int, numero2 int) int {
	return numero1 + numero2
}

// otra manera de establecer los parametros y retornar mas de un dato
func resta(numero1, numero2 int) (int, string) {
	return numero1 - numero2, "String desde la funcion resta"
}

func main() {
	saluda("Eduardo")

	resultado := suma(10, 20)
	fmt.Println("El resultado de sumar 10 + 20:", resultado)

	rest, mensaje := resta(20, 10)
	fmt.Print("El resultado es:", rest, "y el mensaje es:", mensaje)
}
