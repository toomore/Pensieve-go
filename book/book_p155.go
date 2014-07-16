package main

import "fmt"

type Product struct {
    name    string
    price   int
}

//func (product *Product) String() string {
//    return fmt.Sprintf("%s($%d)", (*product).name, (*product).price)
//}

func (product Product) String() string {
    return fmt.Sprintf("%s($%d)", product.name, product.price)
}

func main() {
    cart := []*Product{{"A", 100}, {"B", 200}, {"C", 399}}
    fmt.Println(cart)
    for i := range cart {
        cart[i].price *= 2
    }
    fmt.Println(cart)
}
