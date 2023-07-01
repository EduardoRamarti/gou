package main

import (
	"fmt"
	"os"
)

//recover

/*
La función recover() es una función integrada en Go que se utiliza para capturar y controlar los pánicos (panics) que ocurren durante la ejecución del programa.

Cuando se produce un pánico, normalmente el flujo de ejecución se detiene abruptamente y se inicia una secuencia de desapilamiento de llamadas (stack unwinding) para encontrar una función que pueda manejar el pánico. Si no se encuentra ninguna función que lo maneje, el programa termina su ejecución y muestra un mensaje de error.

La función recover() proporciona una forma de recuperarse de un pánico y controlar el flujo de ejecución. Cuando se llama a recover() dentro de una función defer, se intenta recuperar el valor del pánico y evitar que se propague más allá de ese punto.

*/

func main() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Ups! al parecer el programa no finalizo de forma correcta!")
		}
	}()

	if file, err := os.Open("34_programmingFuncsPt2/example.txt"); err != nil {
		panic("No fue posible obtener el archivo")
	} else {

		defer func() {
			fmt.Println("Cerramos el archivo!!")
			file.Close() //cerramos el archivo
		}()

		contenido := make([]byte, 254)

		longitud, _ := file.Read(contenido)

		contenidoArchivo := string(contenido[0:longitud])

		fmt.Println(contenidoArchivo) //imprimimos el contenido del archivo

	}
}
