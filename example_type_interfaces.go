package main

import "fmt"

type Allprice struct {
    x, y, z float64
}

func (self *Allprice) sum() (total float64) {
    total = self.x + self.y + self.z
    return
}

//http://golang.org/ref/spec#Interface_types
//An interface type specifies a method set called its interface.
type allsum interface {
    sum() float64
}

func main() {
    allprice := Allprice{10, 11, 12}
    fmt.Println(allprice.sum())
    fmt.Println(allsum(&allprice).sum())
}
