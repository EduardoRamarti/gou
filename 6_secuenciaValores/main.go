package main

import "fmt"

// Estructura de const () para declarar la n cantidad de constantes que queramos
const (
	//iota sirve para indicar una secuencia de numeros enteros
	//iota empieza desde 0
	//si queremos iniciar desde otro valor (como 10) basta con colocar iota + 10, y sigue la secuencia
	Domingo int = iota
	Lunes
	Martes
	Miercoles
	Jueves
	Viernes
	Sabado
)

func main() {

	fmt.Println(Domingo)
	fmt.Println(Lunes)
	fmt.Println(Martes)
	fmt.Println(Miercoles)
	fmt.Println(Jueves)
	fmt.Println(Viernes)
	fmt.Println(Sabado)

}
