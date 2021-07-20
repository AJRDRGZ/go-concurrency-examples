package main

import "testing"

func Benchmark_fetch(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fetchSequential(urls)
	}
}

func Benchmark_fetchConcurrent(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fetchConcurrent(urls)
	}
}

func Benchmark_fetchConcurrentCSP(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fetchConcurrentCSP(urls)
	}
}
