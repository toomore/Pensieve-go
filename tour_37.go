//http://tour.golang.org/#37
package main

import "fmt"

var pow = make([]int, 10)

func main() {
    fmt.Println(pow)

    for i := range pow {
        pow[i] = 1 << uint(i)
    }

    for _, v := range pow {
        fmt.Printf("%d\n", v)
    }
}
