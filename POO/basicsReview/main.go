package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	//Varaibles
	var x int
	x = 8
	y := 10
	fmt.Println("Variable x:", x)
	fmt.Println("Variable y:", y)
	//Const
	const u = 67
	fmt.Println("u es una constante con el valor:", u)

	//Parsear valores
	parsedValue, err := strconv.ParseInt("NaN", 0, 64) // (value, base, int size) base = 0 infers automatically the base
	//Errors
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(parsedValue)
	}

	//Maps - Creating dictionaries
	m := make(map[string]int)
	m["Key1"] = 6
	m["Key2"] = 7
	m["Key3"] = 8

	fmt.Println("From map m, Key3 value", m["Key3"])

	//slices
	s := []int{1, 2, 3}
	for i, val := range s {
		//i = index inside slice
		//val = value in that index
		fmt.Println(i, val)
	}

	//Adding new value in our slice
	s = append(s, 8)
	fmt.Println("Elemento agregado al Slice: ", s[3])

	//Channel
	c := make(chan int)

	//Go routine
	go doSomething(c)
	<-c

	//Apuntadores
	g := 25
	fmt.Println("Varialbe g:", g)
	h := &g                                  //usando el apuntador a g
	fmt.Println("Valor en referencia h:", h) //imprime la direccion en memoria
	fmt.Println("Valor alojado en h:", *h)   //el valor en especifico
}

// function
func doSomething(c chan int) {
	time.Sleep(3 * time.Second)
	fmt.Println("Done")
	c <- 1
}
