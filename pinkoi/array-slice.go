package main

import "fmt"

func main() {
    var a [5]int8
    fmt.Println(a)
    b := []int8{1,2,3}
    fmt.Println(b)
    b = append(b, []int8{1,2,3}...)
    fmt.Println(b)
}
