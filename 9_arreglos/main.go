package main

import "fmt"

func main() {

	// Los arreglos son estaticos, entonces deberemos establecer su tamaño desde su declaracion, ademas de incluir el tipo de dato que tendra

	/* Forma 1
	var numeros [5]int

	numeros[0] = 100
	numeros[1] = 200
	numeros[2] = 300
	numeros[3] = 400
	numeros[4] = 500
	*/

	//Forma 2
	numeros := [5]int{100, 200, 300, 400, 500}

	//Forma 3
	numeros2 := [...]int{100, 200, 300, 400, 500}
	// En esta forma, no se coloca su tamaño pero este es determinado por el numero de elementos dentro de el {}

	fmt.Println(numeros)
	fmt.Println(numeros2)

	//Asignar valor a los indices
	//Podemos asignar un valor a un indice en especifico sin necesidad de establecerlos en orden, se ordenaran por si mismos
	monedas := [...]string{3: "Peso Mexicano", 1: "Dolar", 0: "Euro", 2: "Yen", 4: "Soles", 5: "Libra"}
	fmt.Println(monedas)

	//Subarrays (Slices) - obtener un array de otro array
	subMonedas := monedas[0:3] // ["Euro", "Dolar", "Yen"]
	fmt.Println(subMonedas)
}
