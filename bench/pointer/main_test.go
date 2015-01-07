package main

import (
	"fmt"
	"testing"
)

func BenchmarkA(t *testing.B) {
	a := Calcu{Num: 0}
	for i := 0; i < t.N; i++ {
		a.A()
	}
}
func BenchmarkB(t *testing.B) {
	b := Calcu{Num: 0}
	for i := 0; i < t.N; i++ {
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
