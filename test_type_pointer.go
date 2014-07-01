package main

import "fmt"

func main() {
    a := 50
    var b *int
    b = &a
    fmt.Println(a, b)
}
