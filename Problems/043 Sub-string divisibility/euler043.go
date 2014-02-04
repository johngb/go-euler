/*
The number, 1406357289, is a 0 to 9 pandigital number because it is made up of
each of the digits 0 to 9 in some order, but it also has a rather interesting
sub-string divisibility property.

Let d1 be the 1st digit, d2 be the 2nd digit, and so on. In this way, we note
the following:

d2d3d4=406 is divisible by 2
d3d4d5=063 is divisible by 3
d4d5d6=635 is divisible by 5
d5d6d7=357 is divisible by 7
d6d7d8=572 is divisible by 11
d7d8d9=728 is divisible by 13
d8d9d10=289 is divisible by 17

Find the sum of all 0 to 9 pandigital numbers with this property.
*/

package main

import (
	"fmt"
	"time"
)

var p = fmt.Println
var pf = fmt.Printf

// timeTrack is used for basic benchmarking in other functions
func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s \n", name, elapsed)
}

// nextLexicoPerm returns true if there is another lexicographic permutation
// possible, as well as a slice showing the permutation
func nextLexicoPerm(arr []int) (bool, []int) {
	// defer timeTrack(time.Now(), "nextLexicoPerm()") // Timer function

	length := len(arr)
	k, l := -1, -1
	// Find the largest index k such that a[k] < a[k + 1]. If no such index exists, the permutation is the last permutation.
	for i := 0; i < length-1; i++ {
		if arr[i] < arr[i+1] {
			k = i
		}
	}
	if k == -1 {
		// permutation has finished
		return false, arr
	}

	// Find the largest index l such that a[k] < a[l].
	for j := 0; j < length; j++ {
		if arr[k] < arr[j] {
			l = j
		}
	}

	// Swap the value of a[k] with that of a[l].
	arr[k], arr[l] = arr[l], arr[k]

	// reverse the sequence from a[k + 1] up to and including the final element a[n]
	// create a new reversed partial slice from a[k+1] up to a[n]
	arrSegment := make([]int, 0)
	for x := k + 1; x < length; x++ {
		arrSegment = append(arrSegment, arr[x])
	}
	// copy the reversed partial slice back to arr
	for y := 0; y < len(arrSegment); y++ {
		arr[length-y-1] = arrSegment[y]
	}
	return true, arr
}

func digitsToInt(arr []int) int {

	number := 0
	mul := 1

	for i := len(arr) - 1; i >= 0; i-- {
		number += arr[i] * mul
		mul *= 10
	}
	return number
}

// JGB: Using various tricks from http://mathproofs.blogspot.nl/2005/09/number-divisibility-tricks.html
func sumSpecialPandigitals1() int {
	defer timeTrack(time.Now(), "sumSpecialPandigitals1()")

	sum := 0
	set := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	isNext := false
	for {
		isNext, set = nextLexicoPerm(set)
		// if no more permutations, then done
		if isNext == false {
			break
		}
		switch {
		// d2d3d4=406 is divisible by 2
		case set[3]%2 != 0:
			continue
		// d3d4d5=063 is divisible by 3
		case (set[2]+set[3]+set[4])%3 != 0:
			continue
		// d4d5d6=635 is divisible by 5
		case set[5] != 0 && set[5] != 5:
			continue
		// d5d6d7=357 is divisible by 7
		case (set[6]+3*set[5]+2*set[4])%7 != 0:
			continue
		// d6d7d8=572 is divisible by 11
		case (set[7]-set[6]+set[5])%11 != 0:
			continue
		// d7d8d9=728 is divisible by 13
		case (set[8]-3*set[7]-4*set[6])%13 != 0:
			continue
		// d8d9d10=289 is divisible by 17
		case (set[9]+10*set[8]+100*set[7])%17 != 0:
			continue
		// if all the other tests fail, then this must be true
		default:
			sum += digitsToInt(set)
		}
	}
	return sum
}

func sumSpecialPandigitals2() int {
	defer timeTrack(time.Now(), "sumSpecialPandigitals2()")

	sum := 0
	set := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	isNext := false
	for {
		isNext, set = nextLexicoPerm(set)
		// if no more permutations, then done
		if isNext == false {
			break
		}
		switch {
		// d2d3d4=406 is divisible by 2
		case (set[3])%2 != 0:
			continue
		// d3d4d5=063 is divisible by 3
		case (set[4]+set[3]+set[2])%3 != 0:
			continue
		// d4d5d6=635 is divisible by 5
		case (set[5]+10*set[4]+100*set[3])%5 != 0:
			continue
		// d5d6d7=357 is divisible by 7
		case (set[6]+10*set[5]+100*set[4])%7 != 0:
			continue
		// d6d7d8=572 is divisible by 11
		case (set[7]+10*set[6]+100*set[5])%11 != 0:
			continue
		// d7d8d9=728 is divisible by 13
		case (set[8]+10*set[7]+100*set[6])%13 != 0:
			continue
		// d8d9d10=289 is divisible by 17
		case (set[9]+10*set[8]+100*set[7])%17 != 0:
			continue
		// if all the other tests fail, then this must be true
		default:
			sum += digitsToInt(set)
		}
	}
	return sum
}

func main() {
	// set := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// p(digitsToInt(set[0:5]))
	p(sumSpecialPandigitals1())
	p(sumSpecialPandigitals2())

}
