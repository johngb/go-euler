package main

import (
	"fmt"
	"math"
	"time"
)

var p = fmt.Println
var pf = fmt.Printf

// Used to time other functions to compare speed
func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s \n", name, elapsed)
}

func triangleNumber(n int) int {
	return n * (n + 1) / 2
}

// factors1 uses a Brute force approach to determine the factors of n
func factors1(n int) []int {
	// defer timeTrack(time.Now(), "factors1()") // Timer function

	factors := make([]int, 0)
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			factors = append(factors, i)
		}
	}
	return factors
}

func numFactors1(n int) int {
	// defer timeTrack(time.Now(), "numFactors1()") // Timer function

	// there are always at least two factors for any number greater than 1
	if n == 1 {
		return 1
	}
	factors := 2
	i := 2
	limit := int(math.Sqrt(float64(n)))
	// pf("limit = %d\n", limit)
	for ; i <= limit; i++ {
		// pf("i = %d\n", i)
		if n%i == 0 {
			factors += 2
			// pf("factors are: %d and %d\n", i, n/i)
		}
	}
	// pf("limit = %d, i = %d\n", limit, i)
	// i has incremented by now, and we want to see if limit is exactly the square root of n, not a roundnd version
	if (i-1)*(i-1) == n {
		// p("factors are the limit")
		factors--
	}
	// p(factors)
	return factors
}

func firstTriangleNumberWithNFactors(n int) int {
	defer timeTrack(time.Now(), "firstTriangleNumberWithNFactors()") // Timer function

	triangle := 0
	numFactors := 0
	i := 1
	for i = 1; ; i++ {
		triangle = triangleNumber(i)
		// pf("The %dth triangle is %d. Num of factors:", i, triangle)
		numFactors = numFactors1(triangle)
		// pf("%d\n", numFactors)
		if numFactors >= n {

			// pf("Num = %d: The %dth triangle number has %d factors\n", i, triangle, numFactors)
			break
		}
	}
	p(numFactors)
	return triangle
}

func main() {
	p(firstTriangleNumberWithNFactors(500))

}
