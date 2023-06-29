package main

import "fmt"

func main() {
	usuarios := make(map[string]string)
	usuarios["Eduardo"] = "eduardo@rm.com"

	correo := usuarios["Eduardo"]

	if correo != "" {
		fmt.Println(correo)
	} else {
		fmt.Println("No fue posible encontrar el valor ")
	}

	//Cuando se obtiene un valor de un mapa, este nos regresa dos valores. El valor y un booleano
	usuarios["Fulanito"] = ""
	correo2, ok := usuarios["Fulanito"]
	if ok { // si ok contiene un true...
		fmt.Println(correo2)
	} else { // si ok contiene un false
		fmt.Println("No fue posible encontrar el valor ")
	}

	// Otra forma de poder visualizarlo
	if correo3, ok := usuarios["Menganito"]; ok {
		fmt.Println(correo3)
	} else {
		fmt.Println("No fue posible encontrar el valor ")
	}
}
