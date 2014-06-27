package main

import "fmt"

func average(args ...float64) float64 {
    return sum(args...) / float64(len(args))
}

func sum(args ...float64) (total float64) {
    for _, v := range args {
        total += v
    }
    return
}

func max(args ...float64) (max float64) {
    max = args[0]
    for _, v := range args[1:] {
        if v > max {
            max = v
        }
    }
    return
}

func main() {
    float64list := []float64{1, 2, 3, 4, 5}

    //http://golang.org/ref/spec#Passing_arguments_to_..._parameters
    fmt.Println("SUM", sum(float64list...))
    fmt.Println("AVERAGE", average(float64list...))
    fmt.Println("MAX", max(float64list...))
}
