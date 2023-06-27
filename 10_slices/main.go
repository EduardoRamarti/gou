package main

import "fmt"

func main() {

	//Slice
	numeros := []int{1, 2, 3, 4}

	numeros = append(numeros, 5)
	numeros = append(numeros, 6)
	numeros = append(numeros, 7)
	numeros = append(numeros, 8)
	numeros = append(numeros, 9)
	numeros = append(numeros, 10)

	subNumeros := numeros[0:5]

	subNumeros[0] = 100

	fmt.Println(numeros)
	fmt.Println(subNumeros)

	meses := []string{"Enero", "Febrero", "Marzo", "Abril", "Mayo", "Junio", "Julio", "Agosto", "Septiembre"}

	// un puntero
	// una longitud
	// una capacidad

	longitud := len(meses)  //numeros de elementos en el slice
	capacidad := cap(meses) //el tamaño del slice

	//%v me permite decirle al formateador que use el valor de la variable sin importar de que tipo es
	fmt.Printf("La longitud es: %v, La capacidad es: %v\n", longitud, capacidad)

	meses = append(meses, "Octubre") // si la estructura se encuentra en su capacidad maxima, se genera un nuevo arreglo

	longitud = len(meses)
	capacidad = cap(meses)

	fmt.Printf("La longitud es: %v, La capacidad es: %v\n", longitud, capacidad)
}

/*
La diferencia entre un slice y array:

los slice podemos modificar su tamaño añadiendo elementos, pero si generamos un subslice de el, vamos a tener una referencia, no un nuevo conjunto. Una referencia significa que si cambio el valor en el subconjunto o en el conjunto principal, entonces voy a cambiar en todos los arreglos que deriven del arreglo base.

Los arrays se pueden modificar sus elementos establecidos, pero no se puede cambiar el tamaño de este

*/
