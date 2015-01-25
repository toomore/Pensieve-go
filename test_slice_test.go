package main

import "testing"

func BenchmarkA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a()
	}
}

func BenchmarkB(bb *testing.B) {
	for i := 0; i < bb.N; i++ {
		b()
	}
}

//BenchmarkA        100000             13541 ns/op            8192 B/op          1 allocs/op
//BenchmarkB         50000             31619 ns/op           16377 B/op         11 allocs/op
