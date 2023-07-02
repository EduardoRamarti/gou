package main

import "fmt"

/*
Una relación uno a uno se puede implementar mediante la inclusión de una estructura dentro de otra estructura como un campo. Esto implica que cada instancia de la estructura principal tiene una instancia única y exclusiva de la estructura secundaria asociada a ella.
*/

type User struct {
	Name   string
	Email  string
	Active bool
}

type Student struct {
	User User
	Id   string
}

// representa una relación uno a uno donde cada estudiante está vinculado a un único usuario, y podemos acceder a los detalles generales del usuario y a los detalles específicos del estudiante

func main() {

	eduardo := User{Name: "Eduardo", Email: "Eduardo@rm.com", Active: true}
	fulanito := User{Name: "Fulano", Email: "fulano@rm.com", Active: true}

	eduardoStudent := Student{User: eduardo, Id: "1FRTG4"}

	fmt.Println(fulanito)
	fmt.Println(eduardoStudent.User.Name)
	fmt.Println(eduardoStudent.User.Active)
}
