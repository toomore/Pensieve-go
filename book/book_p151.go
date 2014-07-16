package main

import "fmt"

func main() {
    a := make([]int, 5, 10)
    fmt.Println(a, len(a), cap(a))
    a = a[:cap(a)] // len == cap
    fmt.Println(a, len(a), cap(a))

    b := a[:6]
    fmt.Println(b, len(b), cap(b)) // 5, 10
    b = a[2:6]
    fmt.Println(b, len(b), cap(b)) // 4, 8
    b = a[9:]
    fmt.Println(b, len(b), cap(b)) // 1, 1
}
