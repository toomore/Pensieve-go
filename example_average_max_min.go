package main

import "fmt"

func average(args ...float64) float64 {
    var total float64
    for _, v := range args {
        total += v
    }
    return total / float64(len(args))
}

func main() {
    float64list := []float64{1, 2, 3, 4, 5}

    //http://golang.org/ref/spec#Passing_arguments_to_..._parameters
    fmt.Println(average(float64list...))
}
