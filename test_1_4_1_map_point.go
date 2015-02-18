package main

func main() {
	var m = make(map[int]int)
	for i := 0; i < 50000000; i++ {
		m[i] = 1
	}
}
