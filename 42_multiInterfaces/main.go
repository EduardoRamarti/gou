package main

import "fmt"

type Animal interface {
	Dormir()
}

type Mascota interface {
	Comer()
}

//Gato sera una estructura que formara parte de ambas interfaces anteriores, por ello es que tiene un metodo para cada una

type Gato struct {
	//atributes
	Nombre string
}

func (me Gato) Dormir() {
	fmt.Println("El gato duerme")
}

func (me Gato) Comer() {
	fmt.Println("El gato come")
}

func funcionParaUnAnimal(animal Animal) {
	fmt.Println("El objeto es un animal")
}

func funcionParaUnaMascota(mascota Mascota) {
	fmt.Println("El objeto es una mascota")
}

func main() {

	michi := Gato{Nombre: "michi"}

	michi.Comer()
	michi.Dormir()

	funcionParaUnAnimal(michi)
	funcionParaUnaMascota(michi)

}
