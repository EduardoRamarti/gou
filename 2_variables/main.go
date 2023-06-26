package main

import "fmt"

func main() {
	//Forma 1
	//Declarando variables
	var name string // ""
	var age int     // 0

	//Inicializando variables
	name = "Eduardo"
	age = 23

	//Declarando e inicializando variables
	var hobbies string = "Dance"

	//Imprimiendo varaibles
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(hobbies)

	//Forma 2
	//Declarar e inicial variables sin indicar el tipo de dato
	height := 1.67
	weight := 66.2
	//Despues de recibir el tipo de dato que le asignamos, ese sera su tipo de dato por siempre

	//Mostrar variables en pantalla
	fmt.Println(height)
	fmt.Println(weight)

	//Forma 3
	//Declaracion e Inicializacion
	var edad = 23
	var apellido = "RM"

	//Mostrando en pantalla
	fmt.Println(edad)
	fmt.Println(apellido)

	//Otra forma de imprimir en pantalla
	fmt.Println("Mi nombre es:", name, apellido)

	//Declarar multiples variables
	/*
		Forma 1
		var address, zip, pais string = "street 2", "96352", "Mexico"

		Forma 2
		var address, zip, pais = "street 2", "96352", "Mexico"
	*/
	//Forma 3
	address, zip, pais := "street 2", "96352", "Mexico"
	fmt.Println(address, zip, pais)
}
