package main

import "fmt"

type Allprice struct {
    x, y, z float64
}

func (self *Allprice) sum() (total float64) {
    total = self.x + self.y + self.z
    return
}

func (self *Allprice) average() float64 {
    return self.sum() / self.len()
}

func (self *Allprice) len() float64 {
    return float64(3)
}

func name(s string) string {
    return "My name is" + s
}


//http://golang.org/ref/spec#Interface_types
//An interface type specifies a method set called its interface.
type allsum interface {
    sum() float64
}

type allbase interface {
    sum() float64
    average() float64
    len() float64
    //name() string
}

func main() {
    allprice := Allprice{10, 11, 12}
    fmt.Println(allprice.sum())
    fmt.Println(allsum(&allprice).sum())

    allbase := allbase(&allprice)
    fmt.Println(allbase.sum())
    fmt.Println(allbase.average())
}
