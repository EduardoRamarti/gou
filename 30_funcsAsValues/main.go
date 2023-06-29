package main

import "fmt"

type customFunction func(num int) int

func factorial(numero int) int {

	if numero == 1 {
		return 1
	}

	return factorial(numero-1) * numero
}

func main() {

	//una variable llamada miFuncion de tipo func(n int)int
	var miFuncion customFunction
	if miFuncion == nil {
		miFuncion = factorial
	}

	resultado := miFuncion(3)
	fmt.Println(resultado)

}
