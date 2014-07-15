package main

import "regexp"
import "fmt"
import "strings"


func MatchStringSplit(s string, reg *regexp.Regexp) {
    for _, str := range strings.Split(s, ",") {
        str = strings.TrimSpace(str)
        fmt.Println(reg.String(), str, reg.MatchString(str))
    }
}

func main() {
    reg := regexp.MustCompile(`^[a-z0-9\s]+$`)
    fmt.Println(reg.FindAllString("book 123", -1))
    fmt.Println(reg.MatchString("book 123"))

    reg2 := regexp.MustCompile(`^[[:word:]]+$`)
    MatchStringSplit("toomore, toomore0929, 123, tomor!@#, too_more, A B", reg2)
}
