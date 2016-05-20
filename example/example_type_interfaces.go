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

// multi type struct
// `allsum` is a interface
// http://golangtutorials.blogspot.tw/2011/06/interfaces-in-go.html
type alltotalSum struct{
    //allprice []Allprice
    allprice []allsum
}

func (self *alltotalSum) totalSum() (result float64) {
    for _, v := range self.allprice {
        result += v.sum()
    }
    return

}

func main() {
    allprice := Allprice{10, 11, 12}
    fmt.Println(allprice.sum())
    fmt.Println(allsum(&allprice).sum())

    allbase := allbase(&allprice)
    fmt.Println(allbase.sum())
    fmt.Println(allbase.average())

    allprice2 := Allprice{5, 6, 7}
    //multiAttr := [...]allsum{&allprice, &allprice2}
    //fmt.Println(multiAttr)
    multi := alltotalSum{[]allsum{&allprice, &allprice2}}
    fmt.Println(multi)
    fmt.Println(multi.totalSum())
}
