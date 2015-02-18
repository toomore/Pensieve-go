package main

import "fmt"

func main() {
	var m = make(map[int]int)
	for i := 0; i < 100; i++ {
		m[i] = 1
	}
	fmt.Println(m)
}
