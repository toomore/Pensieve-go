package main

import "fmt"

func main() {
    fmt.Println("Find min")
    fmt.Println(Min([]int8{12,41,23,54,123,3,}))
}

func Min(data []int8) (min int8) {
    min = int8(data[0])
    for _, v := range data[1:]{
        if v < min {
            min = v
        }
    }
    return
}
