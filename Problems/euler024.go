/*
A permutation is an ordered arrangement of objects. For example, 3124 is one
possible permutation of the digits 1, 2, 3 and 4. If all of the permutations
are listed numerically or alphabetically, we call it lexicographic order. The
lexicographic permutations of 0, 1 and 2 are:

012   021   102   120   201   210

What is the millionth lexicographic permutation of the digits 0, 1, 2, 3, 4,
5, 6, 7, 8 and 9?
*/

package main

import (
	"fmt"
	"time"
)

var p = fmt.Println
var pf = fmt.Printf

// Used to time other functions to compare speed
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

func repeatLexicoPerm(arr []int, repetitions int) []int {

	for i := repetitions; i > 0; i-- {
		_, arr = nextLexicoPerm(arr)
	}
	return arr
}

func main() {
	// set := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	set := []int{3, 4, 5, 6, 7, 8}

	// set[2], set[3] = set[3], set[2]
	// p(set)

	p(repeatLexicoPerm(set, 999999))

}
