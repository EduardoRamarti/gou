package main

import "fmt"

func main() {
	var num = 1

	// Realiza un bucle infinito
	for {
		fmt.Println(num)
		num++
		if num == 100 {
			break
		}
	}
}
