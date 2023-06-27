package main

import "fmt"

func main() {

	// Un map es una colección de pares clave-valor (diccionario), donde cada clave está asociada a un valor.

	dias := make(map[int]string)
	dias[0] = "Domingo"
	dias[1] = "Lunes"
	dias[2] = "Martes"
	dias[3] = "Miercoles"
	dias[4] = "Jueves"
	dias[5] = "Viernes"
	dias[6] = "Sabado"

	fmt.Println(dias)

}
