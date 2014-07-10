package main

import "fmt"
import "strconv"

func main() {
    fmt.Println("123")
    fmt.Println("\u221E \U0000221a")
    fmt.Println("Z" > "b")
    fmt.Println([]byte("AaBbZz"))
    fmt.Println(strconv.Itoa(123))
    fmt.Println(strconv.FormatUint(123123123, 16))
    fmt.Println([]rune("a國"))
    fmt.Println([]byte("a國"))
}
