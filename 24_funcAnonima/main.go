package main

import "fmt"

func main() {
	/*
		En Go, una función anónima es una función que se define sin un nombre explícito. En lugar de asignarle un nombre, se declara directamente en el lugar donde se va a utilizar, generalmente como una expresión en sí misma.

		Las funciones anónimas en Go se pueden asignar a variables, pasar como argumentos a otras funciones o incluso definirse y llamarse en el mismo lugar.
	*/

	func() {
		fmt.Println("Hola Mundo desde una funcion sin nombre")
	}() //en las funciones anonimas se debe de agregar un parentesis al final para que se llame

	//Guardar una funcion dentro de una variable
	miFunc := func(username string) {
		fmt.Printf("Hola, %s desde una funcion sin nombre\n", username)
	}

	miFunc("Eduardo")

	//Funcion anonima con retorno
	miFuncRec := func(username string) string {
		//se usa sprintf para construir la cadena conformato, pero no la va a imprimir en consola
		message := fmt.Sprintf("Hola, %s desde una funcion anonima con retorno\n", username)
		return message
	}
	miFuncRec("Eduardo")
}
