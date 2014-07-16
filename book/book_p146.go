package main

import "fmt"

func inflate(numbers []int, factor int) {
    fmt.Println(&numbers)
    for i := range numbers {
        numbers[i] *= factor
    }
}

func main() {
    numbers := []int{1,2,3,4,5,6} //array pointer
    fmt.Println(&numbers)
    inflate(numbers, 3)
    fmt.Println(numbers)
}
