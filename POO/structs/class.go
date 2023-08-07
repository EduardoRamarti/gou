package main

import "fmt"

// estructura (equivalente a clase en go)
type Employee struct {
	id   int
	name string
}

// Metodo para la estructura anterior
func (me *Employee) SetId(id int) {
	me.id = id
}

func (me *Employee) SetName(name string) {
	me.name = name
}

func (me *Employee) GetId() int {
	return me.id
}
func (me *Employee) GetName() string {
	return me.name
}

func main() {
	//instancia - new object
	e := Employee{}
	fmt.Printf("%v\n", e)

	e.id = 0
	e.name = "Eduardo"
	fmt.Println("El nombre del empleado es:", e.name)

	e2 := Employee{}
	e2.SetId(4)
	e2.SetName("Anonimo")
	fmt.Println("El nombre del e2:", e2.GetName())
	fmt.Println("El id de e2:", e2.GetId())
}
