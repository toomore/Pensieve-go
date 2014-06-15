//http://tour.golang.org/#25
package main

import "fmt"

func Sqrt(x float64) float64 {
    z := float64(1)
    result_tmp := float64(0)
    result := float64(0)

    for z < x {
        result_tmp = z - (((z * z) - x) / 2 * z)
        if result_tmp < 0 {
           return result
        } else {
            result = result_tmp
        }
        z += 1
    }
    return z
}

func main() {
    fmt.Println(Sqrt(9))
}
