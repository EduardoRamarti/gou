package main

import "fmt"

//Variadic Function
/*
	Es una función que puede recibir un número variable de argumentos del mismo tipo. Esto se logra mediante el uso de los tres puntos suspensivos (...) antes del tipo del último parámetro de la función.

	La principal utilidad de las variadic functions en Go es permitir a los desarrolladores trabajar con un número indeterminado de argumentos sin necesidad de especificar su cantidad exacta.
*/

func promedio(calificaciones ...int) float32 {
	var promedio float32
	var suma int
	for _, calificacion := range calificaciones {
		suma += calificacion
		// fmt.Println(calificacion)
	}
	promedio = float32(suma) / float32(len(calificaciones))
	fmt.Println(promedio)
	return promedio
}

func main() {

	// fmt.Println("Hola", "adios", "Bounjorne")
	rest := promedio(10, 8, 7, 10, 7)
	fmt.Println(rest)
}
