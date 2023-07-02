package main

import "fmt"

type User struct {
	Name  string //atributos
	Email string
}

func (me *User) setName(name string) {
	me.Name = name
}

// func -> definir que tendremos un metodo
// (me *User) -> me es el receptor del metodo, como el python self o java this
//  			 *User indicar hacia donde actuara el metodo, es decir, la estructura User es el propietario y sera el que debera
// setName(name string) -> setName es el nombre del metodo, name el parametro y string el tipo de dato

func (me *User) getName() string {
	return me.Name
}

func main() {

	coder := User{Name: "Hacke", Email: "prub@prub.com"}

	fmt.Println(coder)

	coder.setName("Anonimo")

	fmt.Println(coder)

	fmt.Printf("Retornando el nombre con un metodo: el nombre es %s\n", coder.getName())
}
