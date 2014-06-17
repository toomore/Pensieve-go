//http://tour.golang.org/#37
package main

import "fmt"

var pow = make([]int, 5)

func main() {
    for i := range pow {
        pow[i] = 1 << uint(i)
    }

    fmt.Println(pow)

    for _, v := range pow {
        fmt.Printf("%d\n", v)
    }
}
