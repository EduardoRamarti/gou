package main

import "fmt"

// Creamos una funcion que recibe un entero y retorna un entero
func factorial(numero int) int {

	if numero == 1 { // si el argumento es igual a 1 ...
		return 1 //regresa 1
	}
	//Hace el llamado a si misma, y el retorno de esa llamada se multiplica por el numero
	return factorial(numero-1) * numero
}

func main() {

	resultado := factorial(3)
	fmt.Println("El factorial es:", resultado)

}
