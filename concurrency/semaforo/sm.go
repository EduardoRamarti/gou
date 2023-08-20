package main

import (
	"fmt"
	"sync"
	"time"
)

/*
traffic light.

this uses channels and wait groups to 1. execute only 2 doSmth() func
at a time and 2. be able to wait for all of them.

in order of execution it'll:
c := [][] -- two free spaces
c := [routine][] -- one free space
c := [rountine][routine] -- all occupied
c := [][routine] -- one free space
*/
func main() {
	c := make(chan int, 2) // Creamos un canal con capacidad para dos "instancias" (buffers)
	var wg sync.WaitGroup  // Creamos un objeto WaitGroup

	for i := 0; i < 10; i++ {
		c <- 1               // Asignamos un nuevo "espacio" en el canal
		wg.Add(1)            // Agregamos a la WaitGroup
		go doSmth(i, &wg, c) // Lanzamos una goroutine con la función doSmth y pasamos el índice, la WaitGroup y el canal
	}

	wg.Wait() // Esperamos a que todas las goroutines terminen antes de continuar
}

// Función que simula hacer algo (proceso)
func doSmth(i int, wg *sync.WaitGroup, c chan int) {
	defer wg.Done()                          // Marcamos la tarea como terminada al final de la función
	fmt.Printf("Id: %d -> started...\n", i)  // Imprimimos un mensaje para mostrar que empezamos
	time.Sleep(time.Second * 4)              // Simulamos que toma 4 segundos hacer algo
	fmt.Printf("Id: %d -> finished...\n", i) // Imprimimos un mensaje para mostrar que terminamos

	<-c // Liberamos el espacio para nuevas rutinas
}
