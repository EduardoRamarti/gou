package main

import "fmt"

func main() {
	/*
		&& And
		|| Or
		! Not
	*/

	//Operador And
	resultado := (20 == 20) && (30 == 30) && (20 > 40)
	fmt.Println(resultado)

	//Operador Or
	resultado = (20 == 20) || (30 == 30) || (20 > 40)
	fmt.Println(resultado)

	//Operador Not
	resultado = !(20 == 20)
	fmt.Println(resultado)
}
