package main

import "fmt"

func main() {
    //a := 50
    //var b *int
    //b = &a
    //fmt.Println(a, b)
    var a uint8 = 40
    var b *uint8 = &a
    fmt.Println(a, &a, b)
    *b = 255
    fmt.Println(a, &a, b)
}
