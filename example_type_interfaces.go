package main

import "fmt"

type Allprice struct {
    x, y, z float64
}

func (self *Allprice) sum() (total float64) {
    total = self.x + self.y + self.z
    return
}

func main() {
    allprice := Allprice{10, 11, 12}
    fmt.Println(allprice.sum())
}
