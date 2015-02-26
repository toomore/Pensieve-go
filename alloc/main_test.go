package main

import "testing"

func BenchmarkAlloc(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i <= b.N; i++ {
		Alloc()
	}
}

func BenchmarkAllocaa(b *testing.B) {
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

func BenchmarkAllocee(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		ee()
	}
}

func BenchmarkAllocff(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		ff()
	}
}

func BenchmarkAllocgg(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		gg(i)
	}
}

func BenchmarkAllochh(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		gg(i)
	}
}
