package main

import "fmt"
import "strconv"
import "strings"
import "time"

func main() {
    fmt.Println("123")
    fmt.Println("\u221E \U0000221a")
    fmt.Println("Z" > "b")
    fmt.Println([]byte("AaBbZz"))
    fmt.Println(strconv.Itoa(123))
    fmt.Println(strconv.FormatUint(123123123, 16))
    fmt.Println([]rune("a功"))
    fmt.Println([]byte("a功"))
    tt := time.Now()
    fmt.Println(tt.Nanosecond())

    // string format
    fmt.Printf("%v %v %v \n", "123", 500, "國")
    fmt.Printf("%s %d %T \n", "123", 500, "國")

    // page 108
    fmt.Println(strings.Contains("國家", "國"))
    fmt.Println(strings.Repeat("功蓋許", 10))
    fmt.Println(strings.Title("Toomore is me."))
    fmt.Println(strings.TrimSpace(" I am Toomore! "))
}
