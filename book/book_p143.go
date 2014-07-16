package main

import "fmt"

// Pointer to pointer.
func main() {
    a := 1
    b := &a
    c := &b
    fmt.Println(a, b, c)
    fmt.Println(&a, b, *c)
    fmt.Println(a, *b, **c)
    a++
    *b++
    **c++
    fmt.Println(a, *b, **c)
}
