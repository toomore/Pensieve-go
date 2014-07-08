package main

import "fmt"

type statistics struct {
    numbers []float64
    mean float64
    median float64
}

func Sum(numbers []float64) (sum float64) {
    for _, value := range numbers {
        sum += value
    }
    return
}

func Average(numbers []float64) float64 {
    return Sum(numbers) / float64(len(numbers))
}

func getStatis(numbers []float64) (stats statistics) {
    stats.numbers = numbers
    stats.mean = Average(numbers)
    return
}

func main() {
    numbers := []float64{5, 2, 3, 4, 1}
    fmt.Println(getStatis(numbers))
}
