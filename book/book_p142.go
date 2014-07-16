package main

import "fmt"

// Pointers and values.
func main() {
    a := 10
    b := &a
    fmt.Println(a, *b)
    a++
    fmt.Println(a, *b)
    *b++
    fmt.Println(a, *b)
    c := 100
    b = &c
    fmt.Println(a, *b)
}
