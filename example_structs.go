package main

import "fmt"

type Circle struct {
    x, y, z float64
}

func main() {
    c := new(Circle)
    fmt.Println(c)
}
