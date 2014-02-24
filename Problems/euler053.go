/*
There are exactly ten ways of selecting three from five, 12345:

123, 124, 125, 134, 135, 145, 234, 235, 245, and 345

In combinatorics, we use the notation, 5C3 = 10.

In general,

nCr =    n! / r!(n−r)! ,where r ≤ n, n! = n×(n−1)×...×3×2×1, and 0! = 1. It is
not until n = 23, that a value exceeds one-million: 23 C 10 = 1144066.

How many, not necessarily distinct, values of  nCr, for 1 ≤ n ≤ 100, are
greater than one-million?
*/

// JGB: Not the fastest solution, with lots of room to optimise, but it works
// in 220ms, so good enough

package main

import (
	"fmt"
	"math/big"
	"time"
)

var p = fmt.Println
var pf = fmt.Printf

// timeTrack is used for basic benchmarking in other functions
func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s \n", name, elapsed)
}

func BigFactorial(n int64) *big.Int {
	if n < 0 {
		return big.NewInt(1)
	}
	if n == 0 {
		return big.NewInt(1)
	}
	bigN := big.NewInt(n)
	return bigN.Mul(bigN, BigFactorial(n-1))
}

func nCrGreaterThanMill(n, r int64) bool {
	num := BigFactorial(n)
	den := big.NewInt(1)
	den = den.Mul(BigFactorial(r), BigFactorial(n-r))
	mill := big.NewInt(1e6)
	num = num.Div(num, den)
	if num.Cmp(mill) == 1 {
		return true
	}
	return false
}

func euler053() int {
	defer timeTrack(time.Now(), "euler053()")

	const LIM = 100

	counter := 0
	for r := int64(1); r <= LIM; r++ {
		for n := r; n <= LIM; n++ {
			if nCrGreaterThanMill(n, r) == true {
				counter++
			}
		}
	}
	return counter
}

func main() {

	// p(nCrGreaterThanMill(23, 10))
	p(euler053())

}
