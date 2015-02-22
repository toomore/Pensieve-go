package main

import "testing"

func BenchmarkAlloc(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i <= b.N; i++ {
		Alloc()
	}
}

func BenchmarkAlloc2(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		aa()
	}
}

func BenchmarkAllocbb(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		bb()
	}
}

func BenchmarkAlloccc(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		cc()
	}
}

func BenchmarkAllocdd(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		dd()
	}
}
