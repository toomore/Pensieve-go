package main

import "fmt"

type Person struct {
	name  string
	title []string
	age   int
}

type Author struct {
	name    Person // Named fields
	created string
}

type Author2 struct {
	Person  // Anonymous fields
	created string
}

func exa() {
	points := []struct{ x, y int }{{1, 2}, {3, 4}, {5, 6}} // Anonymous struct.
	for _, point := range points {
		fmt.Printf("%d %d\n", point.x, point.y)
	}
}

func main() {
	exa()
	toomore := Author{Person{"Toomore Chiang", []string{"Python engineer", "Man"}, 30}, "Today"}
	fmt.Println(toomore)
	toomore.name.age = 31
	fmt.Println(toomore)

	fmt.Println("----- Author 2 -----")
	toomore2 := Author2{Person{"Toomore Chiang", []string{"Python engineer", "Man"}, 30}, "Today"}
	fmt.Println(toomore2)
	toomore2.age = 31
	fmt.Println(toomore2)
}
