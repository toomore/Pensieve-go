package main

import "fmt"

func main() {
    var a int
    var b string
    var c float64
    var d byte

    fmt.Printf("%d %s %f %v\n", a, b, c, d)
    if b == "" {
        fmt.Println("b is ''")
    }
}
