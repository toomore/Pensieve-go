//http://tour.golang.org/#20
package main

import "fmt"

func main() {
    sum := 0
    for sum < 100 {
        sum += 1
    }
    fmt.Println(sum)
}
