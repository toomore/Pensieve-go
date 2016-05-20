//http://tour.golang.org/#24
package main

import "fmt"
import "math"

func pow(x, n, lim float64) float64 {
    if v := math.Pow(x, n); v < lim {
        return v
    } else {
        fmt.Printf("%g >= %g\n", v, lim)
    }
    return lim
}

func main() {
    //fmt.Println(
    //    pow(3, 2, 10),
    //    pow(3, 3, 20),
    //)
    fmt.Println(pow(3, 2, 10))
    fmt.Println(pow(3, 3, 20))
}
