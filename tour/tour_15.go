//http://tour.golang.org/#15
package main

import "fmt"
import "math"

func main() {
    //var x, y int = 3, 4
    var x, y = 3, 4
    //var f float64 = math.Sqrt(float64(3*3 + 4*4))
    var f = math.Sqrt(float64(3*3 + 4*4))
    //var z int = int(f)
    var z = int(f)
    fmt.Println(x, y, z)
}
