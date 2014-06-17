//http://tour.golang.org/#35
package main

import "fmt"

func main() {
    var z []int
    fmt.Println(z, len(z), cap(z))
    if z == nil {
        fmt.Println("Is nil!")
    }
}
