package main

import "fmt"

func main() {
    fmt.Println("average")
    x := []int8{1, 2, 3, 4, 5}
    fmt.Println(average(x))
}

func average(list []int8) (total int8) {
    //total = 0
    for _, v := range list {
        total += v
    }
    //return total
    return
}
