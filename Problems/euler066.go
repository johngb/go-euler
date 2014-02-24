/*
Consider quadratic Diophantine equations of the form:

x^2 – D*y^2 = 1

For example, when D=13, the minimal solution in x is 649^2 – 13×180^2 = 1.

It can be assumed that there are no solutions in positive integers when D is
square.

By finding minimal solutions in x for D = {2, 3, 5, 6, 7}, we obtain the
following:

3^2 – 2×2^2 = 1
2^2 – 3×1^2 = 1
9^2 – 5×4^2 = 1
5^2 – 6×2^2 = 1
8^2 – 7×3^2 = 1

Hence, by considering minimal solutions in x for D ≤ 7, the largest x is
obtained when D=5.

Find the value of D ≤ 1000 in minimal solutions of x for which the largest
value of x is obtained.
*/

package main

import (
	"fmt"
	"math"
	"math/big"
	"time"
)

var pl = fmt.Println
var pf = fmt.Printf

// timeTrack is used for basic benchmarking in other functions
func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s \n", name, elapsed)
}

func minimalX(d int) int {
	const LIMIT = 1000000000
	for x := 2; x < LIMIT; x++ {
		y2 := float64(x*x-1) / float64(d) // solution to y^2 from the equation x^2 - D*y^2 = 1
		if isPerfectSquare(y2) {
			return x
		}
	}
	pl("No solution found with a limit of", LIMIT, "with d =", d)
	return 0
}

func isPerfectSquare(potentialSquare float64) bool {

	intRoot := int(math.Sqrt(float64(potentialSquare)))
	// if potentialSquare is a perfect square
	if float64(intRoot*intRoot) == potentialSquare {
		return true
	}
	return false
}

func sqrtInt(n int) int { return int(math.Sqrt(float64(n))) }

func absInt(n int) int {
	if n >= 0 {
		return n
	}
	return -n
}

// finds the minimal X in Pell's equation using the Chakravala method.  Often runs out of digits
func pell(d int) int {
	// Using the Chakravala method

	p, k, x1, y, sd := 1, 1, 1, 0, sqrtInt(d)
	x := 1

	for k != 1 || y == 0 {
		p = k*(p/k+1) - p
		p = p - int((p-sd)/k)*k

		x = (p*x1 + d*y) / absInt(k)
		y = (p*y + x1) / absInt(k)
		k = (p*p - d) / k

		x1 = x
	}
	return x
}

// brute force method that is far too slow to be usable when we get higher numbers
func euler066(lim int) int {
	defer timeTrack(time.Now(), "euler066()")

	maxX := 0
	dWithMaxX := 0
	for d := 2; d <= lim; d++ {

		// if d is a perfect square, skip
		if isPerfectSquare(float64(d)) {
			continue
		}
		// x := minimalX(d)
		x := pell(d)
		if x > maxX {
			maxX = x
			dWithMaxX = d
		}
	}
	pl("Largest minimal x is", maxX, "when d =", dWithMaxX)
	return dWithMaxX
}

// finds the minimal X in Pell's equation using the Chakravala method. Uses math/big for arbitrary precision maths
func bigPell(d int) *big.Int {

	p := big.NewInt(1)
	k := big.NewInt(1)
	x1 := big.NewInt(1)
	y := big.NewInt(0)
	sd := big.NewInt(int64(math.Sqrt(float64(d))))
	x := big.NewInt(1)
	bigD := big.NewInt(int64(d))
	bigZero := big.NewInt(0)
	bigOne := big.NewInt(1)
	t1 := big.NewInt(0)
	t2 := big.NewInt(0)

	// for k != bigOne || y == bigZero {
	for k.Cmp(bigOne) != 0 || y.Cmp(bigZero) == 0 {

		// p = k*(p/k+1) - p
		t1.Div(p, k)
		t1.Add(t1, bigOne)
		t1.Mul(t1, k)
		p.Sub(t1, p)

		// p = p - int((p-sd)/k)*k
		t1.Sub(p, sd)
		t1.Div(t1, k)
		t1.Mul(t1, k)
		p.Sub(p, t1)

		// x = (p*x1 + d*y) / absInt(k)
		t1.Mul(p, x1)
		t2.Mul(bigD, y)
		t1.Add(t1, t2)
		t2.Abs(k)
		x.Div(t1, t2)

		// y = (p*y + x1) / absInt(k)
		t1.Mul(p, y)
		t1.Add(t1, x1)
		y.Div(t1, t2)

		// k = (p*p - d) / k
		t1.Mul(p, p)
		t1.Sub(t1, bigD)
		k.Div(t1, k)

		// x1 = x
		x1.Set(x)
	}
	return x
}

// Fast method that uses bigPell which is based on the Chakravala method.
func bigEuler066(lim int) int {
	defer timeTrack(time.Now(), "bigEuler066()")

	maxX := big.NewInt(0)
	dWithMaxX := 0
	x := big.NewInt(0)

	for d := 2; d <= lim; d++ {

		// if d is a perfect square, skip
		if isPerfectSquare(float64(d)) {
			continue
		}

		x.Set(bigPell(d))
		if x.Cmp(maxX) == 1 {
			maxX.Set(x)
			dWithMaxX = d
		}
	}
	pl("Largest minimal x is", maxX, "when d =", dWithMaxX)
	return dWithMaxX
}

func main() {

	pl(bigEuler066(1000))

}
