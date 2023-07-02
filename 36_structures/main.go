package main

import "fmt"

/*
Una estructura en Go es como un contenedor que nos permite agrupar diferentes tipos de datos relacionados en una sola unidad. Imagina que quieres almacenar información sobre una persona, como su nombre, edad y altura. Una estructura te permite combinar estos datos en un solo objeto.


*/

type User struct {
	//atributos
	Name  string
	Email string
}

func main() {

	//como si se instanciara
	var coder User //es como si fuera un objeto

	//agregando valores a cada atributo
	coder.Name = "Hackerman"
	coder.Email = "info@prub.com"

	fmt.Println(coder)
	fmt.Println(coder.Name)
	fmt.Println(coder.Email)

	//simulando instanciar con un constructor
	coder2 := User{Name: "Hackerm", Email: "info2@prub.com"}

	fmt.Println(coder2)
	fmt.Println(coder2.Name)
	fmt.Println(coder2.Email)

	//En go es posible declarar variables de con el mismo nombre de los atributos, el propio go infiere que se habla de ambos por contexto
	Name := "Hac"
	Email := "prub@purb.com"

	//utilizando el constructor
	coder3 := User{Name, Email}

	fmt.Println(coder3)
	fmt.Println(coder3.Name)
	fmt.Println(coder3.Email)
}
