package main

import "fmt"

//func average(args ...float64) float64 {
//    return sum(args...) / float64(len(args))
//}

func (x *Allprice) sum() (total float64) {
    for _, v := range x.args {
        total += v
    }
    return
}

func maxormin(max_min bool, args ...float64) (result float64) {
    result = args[0]
    for _, v := range args[1:] {
        if max_min {
            if v > result {
                result = v
            }
        } else {
            if v < result {
                result = v
            }
        }
    }
    return
}

func max(args ...float64) float64 {
    return maxormin(true, args...)
}

func min(args ...float64) float64 {
    return maxormin(false, args...)
}

type Allprice struct {
    args []float64
}

func main() {
    float64list := []float64{1, 4, 3, 2, 5}

    //http://golang.org/ref/spec#Passing_arguments_to_..._parameters
    //fmt.Println("SUM", sum(float64list...))
    //fmt.Println("AVERAGE", average(float64list...))
    fmt.Println("MAX", max(float64list...))
    fmt.Println("MIN", min(float64list...))

    x := Allprice{float64list}
    fmt.Println(x.args)
    fmt.Println("SUM", x.sum())
}
