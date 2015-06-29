package main

import (
	"net/url"
	"strings"
	"testing"
)

func BenchmarkDoGet(b *testing.B) {
	baseURL, _ := url.Parse("http://google.com/")
	hostnames := strings.Split("www,docs", ",")
	for i := 0; i <= b.N; i++ {
		DoGet(baseURL, "/", hostnames)
	}
}
