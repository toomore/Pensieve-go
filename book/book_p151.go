package main

import "fmt"

func main() {
    a := make([]int, 5, 10)
    fmt.Println(a, len(a), cap(a))
    a = a[:cap(a)] // len == cap
    fmt.Println(a, len(a), cap(a))
}
