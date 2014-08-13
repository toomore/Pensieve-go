package main

import "time"
import "fmt"

func main() {
	t := time.Now()
	u := t.Unix
	fmt.Println(t.Unix(), u())
	fmt.Println(t.Format(time.RFC3339))
	//value := "Thu, 05/19/11, 10:47PM"
	value := "2014/09/29 21:30:45"
	//layout := "Mon, 01/02/06, 03:04PM"

	// layout format rule
	// http://golang.org/src/pkg/time/format.go#L117
	layout := "2006/01/02 15:04:05"

	pt, _ := time.Parse(layout, value)
	fmt.Println(pt)
}
