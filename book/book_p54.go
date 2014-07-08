package main

import "fmt"
import "strings"

const (
    Read = iota
    Write
    Exec
)

type BigFlag int

const (
    Active BigFlag = 1 << iota
    Send
    Receive
)

var flag = Active | Send

func (flag BigFlag) String() string {
    var flags []string
    if flag&Active == Active {
        flags = append(flags, "Active")
    }
    if flag&Send == Send {
        flags = append(flags, "Send")
    }
    if flag&Receive == Receive {
        flags = append(flags, "Receive")
    }
    if len(flags) > 0 {
        return fmt.Sprintf("%d(%s)", int(flag), strings.Join(flags, "|"))
    }
    return "0()"
}

func main() {
    a := 1 << 1 //0010
    b := 1 << 3 //1000
    fmt.Println(a&b) //0000
    fmt.Println(a|b) //1010

    fmt.Println(Read, Write, Exec)
    fmt.Println(Active)   // 001
    fmt.Println(Send)     // 010
    fmt.Println(Receive)  // 100
    fmt.Println(Send|Receive)  // 110
    fmt.Println(Active|Send|Receive)  // 111
}
