//http://tour.golang.org/#16
package main

import "fmt"

const Pi = 3.14

func main() {
    const World = "世界"
    fmt.Println("Hello", World)
    fmt.Println("Hello", Pi, World)

    const result = true
    //result = false
    fmt.Println("Go rules?", result)
}
