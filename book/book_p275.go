package main

import "fmt"

func exa() {
	points := []struct{ x, y int }{{1, 2}, {3, 4}, {5, 6}} // Anontmous struct.
	for _, point := range points {
		fmt.Printf("%d %d\n", point.x, point.y)
	}
}

func main() {
	exa()
}
