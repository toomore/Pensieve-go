package main

import "testing"

func BenchmarkAlloc(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		Alloc()
	}
}
