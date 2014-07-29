package main

import "fmt"

type Counter struct {
	count int
}

func (counter Counter) Plus() {
	counter.count++
}

func main() {
	var counter Counter
	counter.Plus()
	fmt.Println(counter.count)
}
