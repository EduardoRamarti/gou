package main

import "fmt"

// Función para generar números del 1 al 10 y enviarlos a un canal
func Generator(c chan int) {
	for i := 1; i <= 10; i++ {
		c <- i // Ponemos el número en el canal
	}
	close(c) // Cerramos el canal cuando hemos terminado
}

// Función para recibir números de un canal, duplicarlos y enviarlos a otro canal
func Double(in chan int, out chan int) {
	for value := range in { // Por cada número que llega al canal
		out <- 2 * value // Duplicamos el número y lo ponemos en el nuevo canal
	}
	close(out) // Cerramos el nuevo canal cuando hemos terminado
}

// Función para imprimir números de un canal
func Print(c chan int) {
	for value := range c { // Por cada número que llega al canal
		fmt.Println(value) // Imprimimos el número en la pantalla
	}
}

func main() {
	// Creamos dos canales: uno para los números generados y otro para los números duplicados
	generator := make(chan int)
	doubles := make(chan int)

	// Iniciamos la generación de números en una goroutine
	go Generator(generator)

	// Iniciamos el proceso de duplicación en otra goroutine, usando los números del canal generator
	go Double(generator, doubles)

	// Imprimimos los números duplicados usando la función Print
	Print(doubles)
}
