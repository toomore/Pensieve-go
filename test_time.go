package main

import "time"
import "fmt"

func main() {
	t := time.Now()
	u := t.Unix
	fmt.Println(t.Unix(), u())
	fmt.Println(t.Format("2014/09/29"))
	value := "Thu, 05/19/11, 10:47PM"
	pt, _ := time.Parse(value, value)
	fmt.Println(pt)
}
