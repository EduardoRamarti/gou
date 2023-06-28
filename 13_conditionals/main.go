package main

import "fmt"

func main() {

	var edad int
	fmt.Println("Ingresa tu edad: ")
	fmt.Scanf("%d", &edad)

	if edad >= 18 {
		fmt.Println("El usuario es mayor de edad")
	} else {
		fmt.Println("El usario es menor de edad")
	}

	//Multiples Condicionales (ELSE IF)
	var calificacion int

	fmt.Println("Ingresa tu calificacion:")
	fmt.Scanf("%d", &calificacion)

	if calificacion == 10 {
		fmt.Println("Felicidades, obtuviste una calificacion perfecta")
	} else if calificacion == 8 || calificacion == 9 {
		fmt.Println("Aprobaste la materia")
	} else if calificacion == 6 || calificacion == 7 {
		fmt.Println("Aprobaste la materia, pero necesitas estudiar mas")
	} else if calificacion >= 0 && calificacion <= 5 {
		fmt.Println("No aprobaste la materia")
	} else {
		fmt.Println("Calificacion no valida")
	}

	//Declarar varaibles dentro del if statement
	if nombre, edad := "Coder", 7; nombre == "Coder" { // variables := valores ; comparacion
		fmt.Println("Hola", nombre, "te damos la bienvenida al mundo code")
	} else {
		fmt.Println("Los valores son: ", nombre, edad)
	}
}
