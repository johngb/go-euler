/*
We shall say that an n-digit number is pandigital if it makes use of all the
digits 1 to n exactly once; for example, the 5-digit number, 15234, is 1
through 5 pandigital.

The product 7254 is unusual, as the identity, 39 × 186 = 7254, containing
multiplicand, multiplier, and product is 1 through 9 pandigital.

Find the sum of all products whose multiplicand/multiplier/product identity
can be written as a 1 through 9 pandigital.

HINT: Some products can be obtained in more than one way so be sure to only
include it once in your sum.
*/

package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

var p = fmt.Println
var pf = fmt.Printf

// timeTrack is used for basic benchmarking in other functions
func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s \n", name, elapsed)
}

func findPandigitalProductSums() int {
	defer timeTrack(time.Now(), "findPandigitalProductSums()")

	pandigitalMap := make(map[int]bool)

	for i := 1; i <= 98; i++ {
		for j := 123; j <= 9876; j++ {
			if isPandigital(i, j, i*j) {
				pandigitalMap[i*j] = true
			}
		}
	}

	sum := 0
	for key, _ := range pandigitalMap {
		sum += key
	}
	return sum
}

func isPandigital(a int, b int, c int) bool {
	// defer timeTrack(time.Now(), "isPandigital()")

	str := strconv.Itoa(a) + strconv.Itoa(b) + strconv.Itoa(c)
	if len(str) != 9 {
		return false
	}

	for i := 1; i <= 9; i++ {
		// if the string doesn't contain one of the numbers
		if !strings.Contains(str, strconv.Itoa(i)) {
			return false
		}
	}
	pf("a = %d, b = %d, c = %d\n", a, b, c)
	return true
}

func main() {

	// p(isPandigital(39, 186, 7254))
	p(findPandigitalProductSums())

}
