package main

import "fmt"

/*
En lugar de tener un tipo de dato enum específico, en Go se pueden definir un conjunto de constantes sin tipo que representen los valores posibles de una enumeración. Estas constantes suelen estar agrupadas dentro de un bloque de declaración const y se les asigna un valor incremental utilizando la expresión iota.
*/

type UserType int //Creando el tipo de dato para la enumeracion

const ( // se crea la enumeracion
	// Teacher UserType = 1
	// Student UserType = 2
	Teacher UserType = iota + 1
	Student
)

type User struct { //creamos la clase que contendra el usuario y el tipo, (en el que se espera la enum)
	Username string
	Type     UserType
}

func main() {

	fulano := User{Username: "Fulano", Type: Teacher} //instanciamos el objeto y le pasamos en el constructor, la enumeracion
	eduardo := User{Username: "Eduardo", Type: Student}

	fmt.Println(eduardo)
	fmt.Println(fulano)

	if eduardo.Type == Student {
		fmt.Println("El usuario", eduardo.Username, "es un estudiante")
	}

}
