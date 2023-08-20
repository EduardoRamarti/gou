package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Creamos un objeto WaitGroup para esperar a que todas las goroutines terminen
	var wg sync.WaitGroup

	// Usamos un bucle para crear y lanzar 10 goroutines
	for i := 0; i < 10; i++ {
		wg.Add(1)              // Agregamos una tarea al WaitGroup (contamos una goroutine adicional)
		go doSomething(i, &wg) // Lanzamos una goroutine con la función doSomething y pasamos el índice y el WaitGroup
	}

	wg.Wait() // Esperamos a que todas las goroutines terminen antes de continuar
}

// Función que simula hacer algo (proceso)
func doSomething(i int, wg *sync.WaitGroup) {
	defer wg.Done()               // Marcamos la tarea como terminada al final de la función
	fmt.Printf("Started %d\n", i) // Imprimimos un mensaje para mostrar que empezamos
	time.Sleep(2 * time.Second)   // Simulamos que toma 2 segundos hacer algo
	fmt.Println("End")            // Imprimimos un mensaje para mostrar que terminamos
}
