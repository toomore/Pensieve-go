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

var bb_b = strings.Split("123123123", "2")

func bb() []string {
	return bb_b
}

func cc() []string {
	var a []string
	a = append(a, "Toomore")
	return a
}

func dd() []byte {
	var a []byte
	a = append(a, "Toomore"...)
	return a
}

func ee() map[string]string {
	a := make(map[string]string)
	a["name"] = "Toomore"
	a["age"] = "30"
	return a
}

func main() {
	ee()
}
