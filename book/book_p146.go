package main

import "fmt"

func inflate(numbers []int, factor int) {
    fmt.Println(&numbers)
    for i := range numbers {
        numbers[i] *= factor
    }
}

type oo struct {
    x, y, z int
}

func ooo(o *oo, x, y int) {
    (*o).x += x
    (*o).y += y
}

func main() {
    numbers := []int{1,2,3,4,5,6} //array pointer
    fmt.Println(&numbers)
    inflate(numbers, 3)
    fmt.Println(numbers)

    o := oo{1, 2, 3}
    fmt.Println(o)
    ooo(&o, 2, 1) // ref type.
    fmt.Println(o)
}
