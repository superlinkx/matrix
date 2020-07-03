package main

import (
	"fmt"
	"sync"
	"time"

	"golang.org/x/exp/rand"
)

func main() {
	// Main is used for running the test as a single binary
	// `go test -bench=.` in the repo root is better for benchmarking
	var matrix [10000][10000]int32
	for {
		start := time.Now()
		matrixGoBrrr(&matrix)
		fmt.Println(time.Since(start))
	}
}

// Using pointer because we can't just be passing around values near 1GB in size
func matrixGoBrrr(matrix *[10000][10000]int32) {
	// Setup WaitGroup for outer loop concurrency
	var wg sync.WaitGroup
	wg.Add(10000)

	for outer := 0; outer < 10000; outer++ {
		// Use a go routine for each of the outer iterations
		go func(outer int) {
			defer wg.Done()
			for inner := 0; inner < 10000; inner++ {
				localRand := rand.New(rand.NewSource(uint64(time.Now().UnixNano())))
				matrix[outer][inner] = localRand.Int31()
			}
		}(outer)
	}
	// Make sure all goroutines have completed before this function returns
	wg.Wait()
}
