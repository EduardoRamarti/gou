package main

import "fmt"

func main() {
	var calificacion int

	fmt.Println("Ingresa tu calificacion:")
	fmt.Scanf("%d", &calificacion)

	//Sin condicion
	//Este tipo de switch es mejor usarlo cuando usamos mas comparaciones como en un if
	switch { // true implicito
	//Aqui le decimos que dentro del switch evalue si ...
	case calificacion == 10:
		fmt.Println("Felicidades, obtuviste una calificacion perfecta")
	case calificacion == 8 || calificacion == 9:
		fmt.Println("Aprobaste la materia")
	case calificacion == 6 || calificacion == 7:
		fmt.Println("Aprobaste la materia, pero necesitas estudiar mas")
	case calificacion >= 0 && calificacion <= 5:
		fmt.Println("No aprobaste la materia")
	default:
		fmt.Println("Calificacion no valida")
	}

	// Con condicion
	switch calificacion {
	//Aqui le decimos que dentro del switch evalue si la varaible calificacion matcha con alguno de los valores en cada case
	case 10:
		fmt.Println("Felicidades, obtuviste una calificacion perfecta")
	case 8, 9: //si es 8 o es 9
		fmt.Println("Aprobaste la materia")
	case 6, 7: //si es 6 o es 7
		fmt.Println("Aprobaste la materia, pero necesitas estudiar mas")
	case 1, 2, 3, 4, 5:
		fmt.Println("No aprobaste la materia")
	default:
		fmt.Println("Calificacion no valida")
	}

}
