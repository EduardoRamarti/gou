package main

import "fmt"

func main() {

	var numero = 10

	// Las dos primeras sentencias asegunran que se ejecute almenos una vez
	//ok := true -> se iniicializa la varaible para que sea true
	//ok -> se evalua si es verdadero
	//ok = numero < 10 -> almacena en ok si el numero es menor a 10 (true o false)
	for ok := true; ok; ok = numero < 10 {
		fmt.Println(numero)
		numero++
	}
}
