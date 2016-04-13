package main

import (
	"errors"
	"fmt"
)

type Interface interface {
	echo()
}

type Base struct {
	Name string
}

func (b Base) echo() {
	fmt.Println(b.Name)
}

func (b Base) call() {
	fmt.Println(errors.New("No!"))
}

func run(data Interface) {
	data.echo()
}

func main() {
	s := &Base{Name: "S****y"}
	run(s)
	s.call()
}
