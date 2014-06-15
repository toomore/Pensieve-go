//http://tour.golang.org/#33
package main

import "fmt"

func main() {
    p := []int{2, 3, 5, 7, 11, 13}
    fmt.Println("p==", p)
    fmt.Println("p==[2:4]", p[2:4])
    fmt.Println("p==[:4]", p[:4])
    fmt.Println("p==[3:]", p[3:])
    fmt.Println("p==[:]", p[:])

    for i:=0; i < len(p); i++ {
        fmt.Printf("p[%d] == %d \n", i, p[i])
    }
}
