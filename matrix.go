package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	// Main is used for running the test as a single binary
	// `go test -bench=.` in the repo root is better for benchmarking
	var matrix [10000][10000]int
	for {
		start := time.Now()
		matrixGoBrrr(&matrix)
		fmt.Println(time.Since(start))
	}
}

// Using pointer because we can't just be passing around values near 1GB in size
func matrixGoBrrr(matrix *[10000][10000]int) {
	// Setup WaitGroup for outer loop concurrency
	var wg sync.WaitGroup
	wg.Add(10000)

	for outer := 0; outer < 10000; outer++ {
		// Set up a local random source to prevent goroutines slowing down
		localRand := rand.New(rand.NewSource(time.Now().UnixNano()).(rand.Source64))
		// Use a go routine for each of the outer iterations
		go func(outer int) {
			defer wg.Done()
			for inner := 0; inner < 10000; inner++ {
				matrix[outer][inner] = localRand.Int()
			}
		}(outer)
	}
	// Make sure all goroutines have completed before this function returns
	wg.Wait()
}
