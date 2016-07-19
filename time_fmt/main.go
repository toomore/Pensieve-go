package main

import (
	"fmt"
	"time"
)

// ISO8601 format
const ISO8601 = "2006-01-02T15:04:05-07:00"

func main() {
	t := time.Now()
	fmt.Println(t.Format(ISO8601))
	fmt.Println(t.Format(time.RFC3339))
}
