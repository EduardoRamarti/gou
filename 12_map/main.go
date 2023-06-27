package main

import "fmt"

func main() {

	// Un map es una colección de pares clave-valor (diccionario), donde cada clave está asociada a un valor.

	//Otra gran caracteristica de los mapas, es que estos son desordenados, es decir, no tienen indice

	dias := make(map[int]string)
	dias[0] = "Domingo"
	dias[1] = "Lunes"
	dias[2] = "Martes"
	dias[3] = "Miercoles"
	dias[4] = "Jueves"
	dias[5] = "Viernes"
	dias[6] = "Sabado"

	fmt.Println(dias)

	//Creando un mapa que tiene como clave un string y como valor un array con valores de tipo string
	usuarios := make(map[string][]int)
	usuarios["Eduardo"] = []int{1, 2, 3, 4, 5, 6}
	fmt.Println(usuarios)

	//ITERACION SOBRE MAPAS
	//otra manera de crear un mapa
	users := map[int]string{} //intead of using make
	users[1] = "Usuario-1"
	users[2] = "Usuario-2"
	users[3] = "Usuario-3"
	users[4] = "Usuario-4"
	fmt.Println(users)

	//Recorrer el mapa con un ciclo for
	for key, value := range users {
		fmt.Println(key, value)
	}

}
