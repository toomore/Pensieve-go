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

type CounterB struct {
	count int
}

func (counter *CounterB) Add() {
	counter.count++
	counter.count = counter.count * 10
}

func AllAdd(counter ...MMM) {
	for _, value := range counter {
		value.Add()
	}
}

func main() {
	var countA Counter
	var countB CounterB

	AllAdd(&countA, &countB) // Need pointer
	fmt.Println(countA.count)
	fmt.Println(countB.count)
}
