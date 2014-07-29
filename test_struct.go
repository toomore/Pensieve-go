package main

import "fmt"

type Counter struct {
	count int
}

//func (counter Counter) Plus() { //count still = 0
func (counter *Counter) Plus() { //count++
	//counter.count++
	counter.count += 1
}

func main() {
	var counter Counter
	counter.Plus()
	//(&counter).Plus() // the same
	fmt.Println(counter.count)
}
