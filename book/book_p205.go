package main

import "fmt"

func counter(start int) chan int {
	a := make(chan int)
	go func(i int) {
		for {
			a <- i
			i++
		}
	}(start)
	return a
}

func main() {
	counterA := counter(1)
	counterB := counter(100)
	for i := 0; i < 5; i++ {
		a := <-counterA
		fmt.Printf("(A -> %d, B -> %d) ", a, <-counterB)
	}
	fmt.Println()
}
