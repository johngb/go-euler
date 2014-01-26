package main

import (
	"fmt"
	"math/big"
	"time"
)

var p = fmt.Println
var pf = fmt.Printf
var cache [][]int

// Used to time other functions to compare speed
func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s \n", name, elapsed)
}

// paths1 returns the number of unique paths that are possible traversing a
// 'm' by 'n' lattice (grid).
// Uses recursion, so is just too slow
func paths1(x int, y int) int {
	// defer timeTrack(time.Now(), "paths1()") // Timer function
	if x == 1 && y == 1 {
		return 1 // special case needed to terminate the recursive function
	}
	// normal case when not on edges
	if x > 1 && y > 1 {
		return paths1(x-1, y) + paths1(x, y-1)
	}
	// when on x edge (x = 1).  No need to check for y as if y == 1, function would have returned 1 earlier
	if x == 1 {
		return paths1(x, y-1)
	}
	// when on y edge (y=1). No need for conditional here as this is the only other option
	return paths1(x-1, y)
}

// paths2 is a recursive function to count the number of unique paths in an 'x' by 'y' lattice
// expects there to be a global cache to use
func paths2(x int, y int) int {
	// defer timeTrack(time.Now(), "paths2()") // Timer function

	// if we already have a path calculated and saved
	if cache[x][y] != 0 {
		// p("Getting from cache")
		return cache[x][y]
	}
	var result int
	// normal case when not on edges
	if x > 0 && y > 0 {
		result = paths2(x, y-1) + paths2(x-1, y)
	} else if x == 0 {
		result = paths2(x, y-1)
	} else {
		result = paths2(x-1, y)
	}
	cache[x][y] = result
	// p("writing result to cache")
	return result
}

// countPathsInLattice1 will count the number of unique paths in an 'x' by 'y' lattice
// Uses a global cache ('cache') to speed up the recursion, and calls paths2() as the
// recursive function
func countPathsInLattice1(x int, y int) int {
	// Initialise the cache
	cache = make([][]int, x+1)
	for i := range cache {
		cache[i] = make([]int, y+1)
		for j := range cache[i] {
			cache[i][j] = 0
		}
	}
	cache[0][0] = 1
	return paths2(x, y)
}

func binomialExpansion(n int, k int) int {
	// Useful reference on binomials: http://www.zweigmedia.com/RealWorld/tutstats/bincoeffs.html

	// NUMERATOR
	numerator := 1
	for i := n; i > k; i-- {
		// p("num: ", i)
		numerator *= i
	}
	// DENOMINATOR
	denominator := 1
	for j := 1; j <= n-k; j++ {
		// p("den: ", j)
		denominator *= j
	}

	return numerator / denominator
}

func main() {

	// solves the problem using a single function call to a binomial expansion
	p(big.NewInt(0).Binomial(20+20, 20))

	p(countPathsInLattice1(40, 40))

	p(binomialExpansion(20, 14))

	// p(paths1(16, 16))
	// for i := range cache {
	// 	p(cache[i])
	// }
}
