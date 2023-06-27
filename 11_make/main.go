package main

import "fmt"

func main() {
	/*
		La función make() en Go se utiliza para crear y inicializar ciertos tipos de datos, como slices, maps y channels.

		Slice:
		Para crear un slice, utilizamos make() de la siguiente manera:
		make([]tipoDato, longitud, capacidad).

		Map:
		Para crear un map, utilizamos make() de la siguiente manera:
		make(map[tipoClave]tipoValor).

		Channel:
		Para crear un channel, utilizamos make() de la siguiente manera:
		make(chan tipoDato).

		**Recordar que hay que guardar lo que generemos con make en una variable
	*/

	slice := make([]int, 0, 3)

	slice[0] = 100
	slice[1] = 200
	slice[2] = 300

	slice = append(slice, 400)

	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

}
