//http://tour.golang.org/#15
package main

import "fmt"
import "math"

func main() {
    var x, y int = 3, 4
    var f float64 = math.Sqrt(float64(3*3 + 4*4))
    var z int = int(f)
    fmt.Println(x, y, z)
}
