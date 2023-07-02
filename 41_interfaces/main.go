package main

import "fmt"

/*
Imagina que tienes una serie de objetos en tu programa, como canciones y películas, y quieres poder reproducirlos de alguna manera. En lugar de escribir funciones separadas para cada tipo de objeto, puedes utilizar interfaces para definir un conjunto común de acciones que deben poder realizar esos objetos.

En Go, una interfaz es como un contrato que especifica qué métodos debe tener un objeto para cumplir con esa interfaz. Cualquier objeto que implemente todos los métodos requeridos por una interfaz se considera compatible con esa interfaz.
*/

// creamos la interfaz Animal
type Animal interface { // ejecutara los siguientes metodos
	Comer()
	Dormir()
	Cazar()
}

type Perro struct {
	Nombre string //atributos
}

// Se crean los tres metodos
// como estos metodos no modifican el estado interno de la estructura, no es necesario que se use un puntero como receptor del metodo.
func (me Perro) Comer() {
	fmt.Println("El Perro come")
}

func (me Perro) Dormir() {
	fmt.Println("El Perro duerme")
}

func (me Perro) Cazar() {
	fmt.Println("El Perro caza")
}

// aqui solo creamos una funcion que ejecutara la interfaz, la cual recibe un objeto de tipo Animal, y ejecutara sus metodos en una sola llamada.
func ejecutarAcciones(animal Animal) {
	animal.Comer()
	animal.Dormir()
	animal.Cazar()
}

/*
Una interfaz define un conjunto de métodos que un tipo concreto debe implementar, pero no requiere que haya una función específica para ejecutar la interfaz en sí. En su lugar, puedes utilizar la interfaz como un tipo de datos y trabajar con valores que cumplan con esa interfaz.
*/
func main() {

	floki := Perro{Nombre: "floki"}

	fmt.Println(floki)

	ejecutarAcciones(floki)

}
