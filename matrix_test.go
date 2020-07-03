package main

import (
	"testing"
)

// Test to help see if we're getting more collisions than expected
// This should rarely fail if sufficiently random
func TestMatrixGoBrrr(t *testing.T) {
	var matrix1 [10000][10000]int
	var matrix2 [10000][10000]int
	matrixGoBrrr(&matrix1)
	matrixGoBrrr(&matrix2)
	collisions := 0

	for i := 0; i < 10000; i++ {
		for j := 0; j < 10000; j++ {
			if matrix1[i][j] == matrix2[i][j] {
				collisions++
			}
		}
	}

	if collisions > 100 {
		t.Errorf("More than 100 collisions detected. Got %d, wanted < 100", collisions)
	}
}

// Benchmark our matrix generator
func BenchmarkMatrix(b *testing.B) {
	// run the matrix function b.N times
	for n := 0; n < b.N; n++ {
		var matrix [10000][10000]int
		matrixGoBrrr(&matrix)
	}
}
