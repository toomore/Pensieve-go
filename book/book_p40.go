package main

import "fmt"

type polar struct {
    radius float64
    O      float64
}

type cartesian struct {
    x float64
    y float64
}

func init() {
    fmt.Println("init")
}

func main() {
    var p polar
    var c cartesian
    fmt.Println(p, c)
}
