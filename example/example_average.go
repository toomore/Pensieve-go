package main

import "fmt"

func main() {
    fmt.Println("average")
    x := []int8{1, 2, 3, 4, 5}
    fmt.Println(sum(x))
    fmt.Println(average(x))
}

func sum(list []int8) (total int8) {
    //total = 0
    for _, v := range list {
        total += v
    }
    //return total
    return
}

func average(list []int8) (average int8) {
    total := sum(list)
    average = total / int8(len(list))
    return
}
