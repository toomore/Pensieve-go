package main

import "fmt"

func main() {
    x := make([]int8, 5 ,10)
    x[1] = 1
    fmt.Println(len(x))
    fmt.Println(cap(x))
    fmt.Println(x[6:9])
}
