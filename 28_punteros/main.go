package main

import "fmt"

// se crea esta funcion la cual recibe un argumento de tipo string, el * indica que este argumento debe ser un puntero
func modificarValor(parametro *string) {
	*parametro = "Cambio de valor"
	fmt.Println("Durante la funcion", &parametro)
}

func main() {

	var curso = "Go course"

	fmt.Println("Antes de la funcion", curso)

	//llamamos a la funcion y pasamos el puntero como argumento
	//eso lo que hace es una referencia a la direccion de memoria, por lo que sera capaz de modificar su contenido
	modificarValor(&curso) // referencia

	fmt.Println("Despues de la funcion", curso)
}
