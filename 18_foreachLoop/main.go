package main

import "fmt"

func main() {

	animales := [...]string{"Perro", "Gato", "Pez", "Vaca", "Cerdo"}

	//For tradicional
	for index := 0; index < len(animales); index++ {

		elemento := animales[index]
		fmt.Println(elemento)

	}

	//foreach
	// al usar range este nos regresa dos elementos
	//indice -> lugar del elemento en el array
	//elemento -> el valor del elemento
	for indice, elemento := range animales {
		fmt.Println("El indice es:", indice, "el valor es:", elemento)
	}

	//No utilizando uno de esos elementos
	for _, elemento := range animales {
		fmt.Println("El valor es:", elemento)
	}

	for index, _ := range animales {
		fmt.Println("El indice es:", index)
	}
}
