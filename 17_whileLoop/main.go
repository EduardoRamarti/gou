package main

import "fmt"

func main() {
	numero := 12345
	contador := 0

	//While -> se comporta como un while. Realiza la iteracion mientras numero sea mayor a 0
	for numero > 0 {
		numero /= 10
		contador++ //aumenta en una unidad el valor de contador
	}

	fmt.Println("La cantidad de digitos es: ", contador)
}
