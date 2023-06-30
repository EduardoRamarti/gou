package main

import (
	"fmt"
	"os" //este va a servir para manejar el archivo
)

func main() {

	file, err := os.Open("34_programmingFuncsPt2/example.txt") //abrimos el arcivo
	// file: Es una variable que contendrá un puntero al archivo abierto. Puedes usar este puntero para leer o escribir en el archivo, según sea necesario.
	// err: Es una variable de tipo error que representa cualquier error que pueda ocurrir durante la apertura del archivo. Si la apertura es exitosa, err será nil (un valor nulo)

	if err != nil { // comprobando que si es que hubo algun erro, en cado de haberlo...
		panic("No fue posible obtener el archivo")
	}

	//Con defer estamos diciendo que esta funcion anonima sea llamada antes de terminar el programa
	defer func() {
		fmt.Println("Cerramos el archivo!!")
		file.Close() //cerramos el archivo
	}()

	//creamos el slice con make, este slice tiene capacidad de 254 y se inicializa con todos sus elementos con un 0
	contenido := make([]byte, 254)

	longitud, _ := file.Read(contenido)
	//Read -> sirve para leer el contenido del archivo que abrimos
	//contenido -> es el slice en el que vamos a guardar el contenido obtenido del archivo
	//longitud -> es la cantidad de bytes que se leyeron exitosamente del archivo
	//_ -> se utiliza para regresar un erro en caso de haberlo

	contenidoArchivo := string(contenido[0:longitud])
	//contenidoArchivo guardara un string
	//convertimos el contenido en un string

	fmt.Println(contenidoArchivo) //imprimimos el contenido del archivo

}
