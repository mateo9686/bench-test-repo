package main

import (
	"testing"
)

var num = 1000

// BenchmarkPrimeNumbers
func BenchmarkPrimeNumbers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		primeNumbers(num)
	}
	b.Error("bla")
}

func BenchmarkPrimeNumbers2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		primeNumbers(num)
	}
}
