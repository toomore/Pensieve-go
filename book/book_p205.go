package main

import "fmt"
import "time"
import "math/rand"

func counter(start int) chan int {
	a := make(chan int, 5)
	go func(i int) {
		for {
			time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
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
		b := <-counterA
		fmt.Printf("(A -> %d, B -> %d) ", a, <-counterB)
		fmt.Println(b)
	}
	fmt.Println()
}
