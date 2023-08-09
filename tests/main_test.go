package main

import "testing"

func TestSum(t *testing.T) {
	// total := Sum(5, 5)
	// if total != 10 {
	// 	t.Errorf("Sum was incorrect, got %d expected %d", total, 10)
	// }

	// Declaración de una slice de estructuras anónimas con campos a, b y c de tipo int
	tables := []struct {
		a int //parametro1
		b int //parametro2
		c int //Resultado
	}{
		// Primer estructura con valores {1, 2, 3}
		{1, 2, 3},
		// Segunda estructura con valores {2, 2, 4}
		{2, 2, 4},
		// Tercera estructura con valores {25, 26, 51}
		{25, 26, 51},
	}

	for _, item := range tables {
		total := Sum(item.a, item.b)
		if total != item.c {
			t.Errorf("Sum was incorrect, got %d expected %d", total, item.c)
		}
	}
}

func TestMax(t *testing.T) {
	tables := []struct {
		a int //parametro1
		b int //parametro2
		h int //resultado
	}{
		{4, 2, 4},
		{3, 2, 3},
		{1, 2, 2},
	}

	for _, item := range tables {
		max := GetMax(item.a, item.b)
		if max != item.h {
			t.Errorf("Max was incorrect, got %d expected %d", max, item.h)
		}
	}
}

func TestFib(t *testing.T) {
	table := []struct {
		a int //parametro
		n int //resultado
	}{
		{1, 1},            //Caso1
		{8, 21},           //Caso2
		{50, 12586269025}, //Caso3
	}
	for _, item := range table {
		fib := Fibonacci(item.a)
		if fib != item.n {
			t.Errorf("Fibonacci was incorrect, got %d expected %d", fib, item.n)
		}
	}
}
