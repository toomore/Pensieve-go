package main

import "fmt"

// Point interface
type Point interface {
	Run()
}

// Car struct
type Car struct {
	Name string
}

// Run for Car
func (c Car) Run() {
	fmt.Println(c.Name)
}

// RunInterface ...
func RunInterface(p Point) {
	p.Run()
}

func main() {
	var a int64
	b := &a
	fmt.Println(&a, b, &b)

	car := &Car{Name: "Toomore"}
	RunInterface(car)
}
