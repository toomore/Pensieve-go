//http://tour.golang.org/#36
package main

import "fmt"

var pow = []int{1, 2, 3, 5, 8, 11, 13}

func main() {
    fmt.Println(pow)
    for i, v := range pow {
        fmt.Printf("2**%d = %d\n", i, v)
    }
}
