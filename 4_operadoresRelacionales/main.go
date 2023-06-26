package main

import "fmt"

func main() {
	/*
		> mayor que
		< menor que
		>= mayor igual que
		<= menor igual que
		== igual que
		!= diferente a
	*/

	var edad = 10
	var resultado = edad > 5 // devuelve un booleano

	var resultado2 = edad <= 10

	var name = "Gopher"
	var homonimos = name == "Gopher"

	fmt.Println(resultado, resultado2)
	fmt.Println("Son homonimos:", homonimos)
}
