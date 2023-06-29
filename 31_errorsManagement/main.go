package main

import (
	"errors" //importamos el modulo errors
	"fmt"
)

// creamos la funcion, recibe 2 enteros y retorna un entero y un error
func division(dividendo, divisor int) (int, error) {

	if divisor == 0 { //si divisor es igual a cero, entonces...
		return 0, errors.New("no se puede dividir por cero") //regresa cero como valor y crea un nuevo error con un mensaje
	} else {
		return dividendo / divisor, nil //de lo contrario regresa el valor de la division y un nil
	}
}

func main() {

	if resultado, err := division(10, 0); err != nil {
		// fmt.Println(err)
		panic(err)
	} else {
		fmt.Println("El resultado es:", resultado)
	}

}
