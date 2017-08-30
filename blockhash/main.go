package main

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

func check(sbody string, n int, timestamp int) {
	h := sha256.New()
	body := []byte(sbody)
	data := append(body[:], []byte(strconv.Itoa(n))...)
	data = append(data[:], []byte(strconv.Itoa(timestamp))...)
	h.Write(data)
	s := h.Sum(nil)
	fmt.Printf("%x\n", s)
	fmt.Println(data)
}

func checkRepeat(h []byte, n int, b byte) bool {
	for i := 0; i < n; i++ {
		if h[i] != b {
			return false
		}
	}
	return true
}

func work(body []byte, nowStr string, rn int, b byte, n int, done *bool, limit chan struct{}) {
	data := append(body[:], []byte(strconv.Itoa(n))...)
	data = append(data[:], []byte(nowStr)...)
	h := sha256.New()
	h.Write(data)
	s := h.Sum(nil)
	if checkRepeat(s, rn, b) {
		fmt.Println("OK")
		fmt.Printf("%x\n", s)
		fmt.Println("t:", nowStr)
		fmt.Println("n:", n)
		fmt.Println(data)
		*done = false
	}
	<-limit
}

func mine(rn int, b byte) {
	body := []byte("Toomore")
	var n int
	now := time.Now()
	nowStr := strconv.Itoa(int(now.Unix()))
	var done = true
	limit := make(chan struct{}, 16)

	for done {
		limit <- struct{}{}
		go work(body, nowStr, rn, b, n, &done, limit)
		n++
	}
	fmt.Println(time.Now().Sub(now))
}

func main() {
	mine(4, 17)
	//check("Toomore", 47380349, 1504096080)
}
