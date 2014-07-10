package main

import "fmt"
import "strconv"
import "time"

func main() {
    fmt.Println("123")
    fmt.Println("\u221E \U0000221a")
    fmt.Println("Z" > "b")
    fmt.Println([]byte("AaBbZz"))
    fmt.Println(strconv.Itoa(123))
    fmt.Println(strconv.FormatUint(123123123, 16))
    fmt.Println([]rune("a國"))
    fmt.Println([]byte("a國"))
    tt := time.Now()
    fmt.Println(tt.Nanosecond())

    // string format
    fmt.Printf("%v %v %v \n", "123", 500, "國")
    fmt.Printf("%s %d %T \n", "123", 500, "國")
}
