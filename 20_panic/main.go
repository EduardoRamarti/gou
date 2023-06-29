package main

import "fmt"

func main() {
	//panic
	/*
		La función panic se utiliza para detener inmediatamente la ejecución normal de un programa y generar un pánico. Un pánico es una situación excepcional que indica un error grave y generalmente no recuperable.

		Cuando se llama a la función panic, el programa se detiene y se imprime un mensaje de error. Además, se desenrollan (se deshacen en orden inverso) todas las funciones en la pila de llamadas, liberando los recursos en el camino. Esto significa que cualquier código diferido (defer) se ejecutará antes de salir del programa.
	*/

	var dividendo, divisor int

	fmt.Println("Ingresa el valor del dividendo:")
	fmt.Scanf("%v", &dividendo)

	fmt.Println("Ingresa el valor del divisor:")
	fmt.Scanf("%v", &divisor)

	if divisor == 0 {
		panic("No es posible dividir entre 0")
	}

	resultado := dividendo / divisor
	fmt.Println(resultado)

}
