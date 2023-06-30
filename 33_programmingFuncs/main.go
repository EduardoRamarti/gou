package main

import "fmt"

func funcion1() {
	fmt.Println("Soy la func 1")
}

func funcion2() {
	fmt.Println("Soy la func 2")
}

func main() {

	//defer
	/*
		"defer" se utiliza para posponer la ejecución de una función hasta que la función que la rodea haya terminado. Es una forma de programar una tarea para que se ejecute al finalizar una función, sin importar cómo se haya salido de ella (ya sea por un retorno normal o una excepción
	*/

	fmt.Println("Estamos en el bloque main ")
	defer funcion1() //Esta funcion se ejecuta cuando el bloque main finaliza
	funcion2()

}
