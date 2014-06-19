package main

import "fmt"

func main() {
    x := make([]int8, 5 ,10)
    //var x []int
    fmt.Println(len(x))
    fmt.Println(cap(x))
    //fmt.Println(x[6:9])
    y := append(x, 1)
    fmt.Println(y)
    for i, v := range x {
        fmt.Println(i, v)
    }
}
