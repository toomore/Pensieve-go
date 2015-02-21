package main

import "strings"

func Alloc() int {
	var a int
	a = 1
	return a
}

func aa() []string {
	var a string
	a = "123123123"
	var b []string
	b = strings.Split(a, "2")
	return b
}

func main() {
	aa()
}
