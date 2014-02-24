/*
If p is the perimeter of a right angle triangle with integral length sides,
{a,b,c}, there are exactly three solutions for p = 120.

{20,48,52}, 24,45,51}, 30,40,50}

For which value of p ≤ 1000, is the number of solutions maximised?
*/

package main

import (
	"fmt"
	"math"
	"time"
)

var p = fmt.Println
var pf = fmt.Printf

// timeTrack is used for basic benchmarking in other functions
func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s \n", name, elapsed)
}

// isIntegerHyp returns true and the hypotenuse length if the hypotenuse
// corresponding to the right angled triangle with lengths 'a' and 'b'
func isIntegerHyp(a, b int) (bool, int) {
	// defer timeTrack(time.Now(), "isIntegerHyp")

	hyp := int(math.Sqrt(float64(a*a + b*b)))
	if hyp*hyp == a*a+b*b {
		return true, hyp
	}
	return false, -1
}

func findPerimWithMostIntSolutions(lim int) (int, int) {
	defer timeTrack(time.Now(), "findPerimWithMostIntSolutions()")

	perimeterMap := make(map[int]int)

	isHypInt := false
	c := 0
	p := 0
	// Let a always be less than b
	for a := 1; a < lim/3; a++ {
		for b := a + 1; b < lim/2; b++ {
			isHypInt, c = isIntegerHyp(a, b)
			if isHypInt {
				p = a + b + c
				perimeterMap[p] += 1

			}
		}
	}

	// find the greatest value in perimeterMap
	pWithMax := 0
	pSolutions := 0
	for k, v := range perimeterMap {
		if v > pSolutions {
			pWithMax = k
			pSolutions = v
		}

	}
	return pWithMax, pSolutions
}

func main() {
	// p(isIntegerHyp(3, 5))
	p(findPerimWithMostIntSolutions(1000))

}
