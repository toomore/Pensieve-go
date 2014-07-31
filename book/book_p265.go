package main

import "fmt"

type MMM interface {
	Add()
}

type Counter struct {
	count int
}

func (counter *Counter) Add() {
	counter.count++
}

func AllAdd(counter MMM) {
	counter.Add()
}

func main() {
	var countA Counter
	AllAdd(&countA) // Need pointer
	fmt.Println(countA.count)
}
