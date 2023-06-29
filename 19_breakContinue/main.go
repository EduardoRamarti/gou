package main

import "fmt"

func main() {

	//break - continue

	for i := 1; i <= 10; i++ {

		//Cuando 1 sea 5, la instruccion continue hara que ya no se ejecuten las lineas de condigo siguientes y se comience con la siguiente iteracion
		if i == 5 {
			continue
		}

		//Cuando i sea 8, la instruccion break hara que se detenga el ciclo y termine de inmediato, sin que las siguientes lineas de codigo se ejecuten
		if i == 8 {
			break
		}

		fmt.Println(i)
	}

}
