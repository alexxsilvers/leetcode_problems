package main

import "testing"

func BenchmarkFindMissing(b *testing.B) {
	for n := 0; n < b.N; n++ {
		FindMissing([]int{5, 6, 8, 12}, []int{6,8, 12})
	}
}

func BenchmarkFindMissingXOR(b *testing.B) {
	for n := 0; n < b.N; n++ {
		FindMissingXOR([]int{5, 6, 8, 12}, []int{6,8, 12})
	}
}