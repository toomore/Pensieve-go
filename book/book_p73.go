package main

import "fmt"
import "math"
import "sort"

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

func Median(numbers []float64) float64 {
    sort.Float64s(numbers)
    return numbers[int(math.Ceil(float64(len(numbers)) / 2))-1]
}

func getStatis(numbers []float64) (stats statistics) {
    stats.numbers = numbers
    sort.Float64s(stats.numbers)
    stats.mean = Average(numbers)
    stats.median = Median(numbers)
    return
}

func main() {
    numbers := []float64{5, 2, 3, 4, 1}
    fmt.Println(getStatis(numbers))
}
