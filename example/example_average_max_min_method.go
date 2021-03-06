package main

import "fmt"

type Allprice struct {
    args []float64
}

func (x *Allprice) average() float64 {
    return x.sum() / x.len()
}

func (x *Allprice) sum() (total float64) {
    for _, v := range x.args {
        total += v
    }
    return
}

func (x *Allprice) max() float64 {
    return maxormin(true, x.args...)
}

func (x *Allprice) min() float64 {
    return maxormin(false, x.args...)
}

func (x *Allprice) len() float64 {
    return float64(len(x.args))
}

//http://golang.org/ref/spec#Passing_arguments_to_..._parameters
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

func main() {
    float64list := []float64{1, 4, 3, 2, 5}

    x := Allprice{float64list}
    fmt.Println(x.args)
    fmt.Println("LEN", x.len())
    fmt.Println("SUM", x.sum())
    fmt.Println("AVERAGE", x.average())
    fmt.Println("MAX", x.max())
    fmt.Println("MIN", x.min())
}
