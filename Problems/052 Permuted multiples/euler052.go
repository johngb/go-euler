/*
It can be seen that the number, 125874, and its double, 251748, contain
exactly the same digits, but in a different order.

Find the smallest positive integer, x, such that 2x, 3x, 4x, 5x, and 6x,
contain the same digits.
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

// sameDigits returns true if a and b have the same digits in any order
func sameDigits(a, b int) bool {
	strA := strconv.Itoa(a)
	strB := strconv.Itoa(b)

	// if they aren't the same length, the answer can't be True
	if len(strA) != len(strB) {
		return false
	}

	for i := 0; i < len(strA); i++ {
		idx := strings.Index(strB, strA[i:i+1])
		// p("idx =", idx)
		if idx == -1 {
			return false
		} else {
			// p("i =", i, ",char to remove: ", strA[i:i+1])
			strB = strings.Replace(strB, strA[i:i+1], "", 1)
		}
	}
	return true
}

func euler052(lim int) int {
	defer timeTrack(time.Now(), "euler052()")

	for i := 1; i < lim; i++ {

		// if there are to be the same number of digits, the first digit of
		// the answer has to be 1.  If it were 2, 6*2 would increase the
		// number of digits.
		str := strconv.Itoa(i)
		if str[0:1] != "1" {
			i *= 5
		}

		switch {
		case !sameDigits(i, 6*i):
			continue
		case !sameDigits(i, 5*i):
			continue
		case !sameDigits(i, 4*i):
			continue
		case !sameDigits(i, 3*i):
			continue
		case !sameDigits(i, 2*i):
			continue
		default:
			return i
		}
	}
	return -1
}

func main() {
	// p(sameDigits(12113, 32111))

	p(euler052(1000000))
}
