package main

/*
Unbuffered channel: Espera una función o una rutina para recibir el mensaje, es bloqueada por ser llamada en la misma función. Es como pasarse una pelota directamente, tienes que esperar a que el receptor esté listo para recibir antes de continuar.

Buffered channel: Se puede llamar de manera inmediata, en el siguiente ejemplo 2 es el numero de canales que pueden ser usados. Es como darle una pelota a alguien y que la sostenga por un rato antes de usarla. Puedes continuar haciendo otras cosas mientras esperas.
*/

import "fmt"

func main() {
	// c := make(chan int) // Unbuffered
	c := make(chan int, 2) // Buffered

	c <- 1 //enviando 1
	c <- 2 //enviando 2

	fmt.Println(<-c) //imprimiendo lo que trae el canal
	fmt.Println(<-c)
}
