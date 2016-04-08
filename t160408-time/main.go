package main

import (
	"fmt"
	"time"
)

var now = time.Now()

func main() {
	fmt.Println(now.UnixNano())
	fmt.Println(time.Since(now))
	//jsonNow, _ := now.MarshalJSON()
	fmt.Println(now.ISOWeek())
}
