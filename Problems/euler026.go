/*
A unit fraction contains 1 in the numerator. The decimal representation of the
unit fractions with denominators 2 to 10 are given:

1/2 =   0.5 1/3 =   0.(3) 1/4 =   0.25 1/5 =   0.2 1/6 =   0.1(6) 1/7 =
0.(142857) 1/8 =   0.125 1/9 =   0.(1) 1/10    =   0.1 Where 0.1(6) means
0.166666..., and has a 1-digit recurring cycle. It can be seen that 1/7 has a
6-digit recurring cycle.

Find the value of d < 1000 for which 1/d contains the longest recurring cycle
in its decimal fraction part.
*/

/* NOTES from vbay03

A full reptend prime in base b is a prime such that b is a primitive root modulo p.  In this case 10
is the base.

A number m is a primitive root modulo n, if the multiplicative order of the number m modulo n is
equal to phi(n).  (totient function)  (Primitive roots modulo n)

Phi(p) = p - 1.  Therefore the order of 10 modulo p must be a divisor of p - 1  ( ord[p](10)|p-1 )
(Multiplicative order).

10 is a primitive root mod p only if the order is equal to phi(p) or p - 1.  Therefore, for each
proper divisor of p - 1,  q1,q2,q3...  where q[i]<p-1,  if 10^q[i]=1 mod p for any one of them, then
10 is not a primitive root modulo p.

Other References: http://math.stackexchange.com/questions/74884/quick-algorithm-for-computing-
orders-mod-n http://rosettacode.org/wiki/Multiplicative_order
*/

package main

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

var p = fmt.Println
var pf = fmt.Printf

// timeTrack is used for basic benchmarking in other functions
func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s \n", name, elapsed)
}

func inSlice(slice []int, value int) (bool, []int) {
	found := false
	positions := []int{}
	for i := 0; i < len(slice); i++ {
		if slice[i] == value {
			found = true
			positions = append(positions, i)
		}
	}
	return found, positions
}

func repeatingDecimalLength(divisor int) int {
	// defer timeTrack(time.Now(), "repeatingDecimalLength()")

	dividend := int(math.Pow(10, float64(len(strconv.Itoa(divisor)))))
	limit := dividend
	remainderList := make([]int, 0)

	for len(remainderList) <= limit {
		ans := dividend / divisor
		diff := dividend - ans*divisor
		if diff == 0 {
			return 0
		}
		isInRemainderList, position := inSlice(remainderList, diff)
		// p("ans: ", ans, ", diff: ", diff)

		if isInRemainderList {
			// Compare current position to position in list to get the repeating length
			// p(remainderList)
			return len(remainderList) - position[0]
		}
		remainderList = append(remainderList, diff)
		dividend = diff * 10
	}
	return -1 // error condidion
}

func longestRepeatingDecimal(lim int) int {
	defer timeTrack(time.Now(), "longestRepeatingDecimal()")

	longestLength, numberWithLongestRepeat := 0, 0
	repeatLength := 0

	for i := 1; i < lim; i++ {
		repeatLength = repeatingDecimalLength(i)
		if repeatLength > numberWithLongestRepeat {
			numberWithLongestRepeat = i
			longestLength = repeatLength
		}
	}
	p(longestLength)
	return numberWithLongestRepeat
}

// JGB: Could speed up by only testing primes
func main() {

	p(longestRepeatingDecimal(1000))

}
