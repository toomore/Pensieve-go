package main

import (
	"log"
	"strconv"
	"strings"
)

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
	a["address"] = "Taiwan"
	return a
}

func ff() map[string]string {
	a := map[string]string{
		"name": "Toomore",
		"age":  "30",
	}
	a["address"] = "Taiwan"
	return a
}

var gg_a = map[string]string{
	"name": "Toomore",
	"age":  "30",
}

func gg(i int) map[string]string {
	gg_a["address"] = "Taiwan"
	ii := string(i)
	gg_a[ii] = ii
	return gg_a
}

type hhmap map[string]string

func hh(i int) map[string]string {
	hh_a := hhmap{}
	hh_a["address"] = "Taiwan"
	ii := strconv.Itoa(i)
	for range []int{1, 2, 3, 4, 5} {
		hh_a[ii] = ii
	}
	return hh_a
}

func main() {
	log.Println(hh(1))
}
