package main

import (
	"fmt"
	"testing"
)

func BenchmarkA(*testing.B) {
	a := Calcu{Num: 0}
	for i := 0; i < 1000000; i++ {
		a.A()
	}
}
func BenchmarkB(*testing.B) {
	b := Calcu{Num: 0}
	for i := 0; i < 1000000; i++ {
		b.B()
	}
}

func TestA(*testing.T) {
	a := Calcu{Num: 0}
	a.A()
	a.A()
	fmt.Println(a.Num)
}
func TestB(*testing.T) {
	b := Calcu{Num: 0}
	b.B()
	b.B()
	fmt.Println(b.Num)
}
