package main

import "fmt"

var username string //variable global

func funcion1() {
	username = "user1"
	fmt.Println("Funcion1:", username)
}

func funcion2() {
	username = "user2"
	fmt.Println("Funcion2:", username)
}

func main() {

	username = "Boketto"
	funcion1()
	funcion2()

	//Como estamos accediendo a una variable globlar. en el siguient print no habra un "boketto", habra la ultima modificacion que seria la llamada de la funcion 2
	fmt.Println(username)

}
