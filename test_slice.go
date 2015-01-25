package main

import "fmt"

func a() []int {
	a := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		a[i] = 1
	}
	return a
}

func b() []int {
	var a []int
	for i := 0; i < 1000; i++ {
		a = append(a, 1)
	}
	return a
}

//BenchmarkA        100000             13541 ns/op            8192 B/op          1 allocs/op
//BenchmarkB         50000             31619 ns/op           16377 B/op         11 allocs/op

func main() {
	fmt.Println(1)
	fmt.Println(len(b()), len(a()))
}
