package main

import "fmt"

func fibonacci() func(int) int {
    a := 0
    b := 1
    return func(x int) (result int) {
        if x >= 2 {
            result = a + b
            a, b = b, result
        }else if x == 1 {
            result = 1
        }
        return
    }
}

func main() {
    fi := fibonacci()
    for i := 0; i < 10; i++ {
        fmt.Println(i, fi(i))
    }
}

//0: 0
//1: 1
//2: 0 + 1
//3: 1 + 0 + 1
//4: 0 + 1 + 1 + 0 + 1
