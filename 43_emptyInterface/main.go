package main

/*
son interfaces que no especifican ningún método. Esto significa que cualquier tipo en Go implementa automáticamente esta interfaz vacía, ya que no hay requisitos específicos de implementación.

Las interfaces vacías se utilizan en Go para crear funciones o estructuras que puedan manejar valores de diferentes tipos sin preocuparse por el tipo específico.


-> Al utilizar una interfaz vacía como tipo de parámetro, se puede aceptar cualquier tipo de valor como argumento. <-
*/

import "fmt"

//define una función llamada mostrarVariable que toma un parámetro llamado objeto de tipo interface{}.
//El tipo interface{} es una interfaz vacía en Go que puede contener valores de cualquier tipo.

func mostrarVariable(objeto interface{}) {
	fmt.Printf("El valor de la variable es: %v\n", objeto)
}

func suma(n1, n2 int) int {
	return n1 + n2
}

type User struct {
	Name string
}

func main() {
	usuario := User{Name: "Eduardo"}
	mostrarVariable(usuario)
	mostrarVariable(suma)
}
