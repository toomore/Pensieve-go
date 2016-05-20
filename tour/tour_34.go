//http://tour.golang.org/#34
package main

import "fmt"

func printSlice(s string, x []int) {
    fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

func main() {
    a := make([]int, 5)
    printSlice("a", a)
    b := make([]int, 0, 5)
    printSlice("b", b)
    c := b[:2]
    printSlice("c", c)
    d := c[2:5]
    printSlice("d", d)
    e := b[:cap(b)]
    e = e[1:]
    printSlice("e", e)
    //f := b[1:]
    //printSlice("f", f)  //slice bounds out of range
}
