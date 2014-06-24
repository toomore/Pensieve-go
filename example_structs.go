package main

import "fmt"

type Circle struct {
    x, y, z float64
}

func total(x *Circle) float64 {
    return x.x + x.y + x.z
}

func main() {
    //c := new(Circle)
    c := Circle{1, 2, 3}

    fmt.Println(c)
    fmt.Println(total(&c))
}
