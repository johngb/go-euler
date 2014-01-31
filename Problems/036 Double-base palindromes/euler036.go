/*
The decimal number, 585 = 10010010012 (binary), is palindromic in both bases.

Find the sum of all numbers, less than one million, which are palindromic in
base 10 and base 2.

(Please note that the palindromic number, in either base, may not include
(leading zeros.)
*/

package main

import (
	"fmt"
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

// isPalindrome returns true if a string 's' is a palindrome and false otherwise
func isPalindrome(s string) bool {
	// defer timeTrack(time.Now(), "isPalindrome()")

	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

// findSumOfDualPalindromes returns the sum of all the numbers below 'lim'
// which are palindromes in both base 2 and base 10
func findSumOfDualPalindromes(lim int) int {
	defer timeTrack(time.Now(), "findSumOfDualPalindromes()")

	sum := 0
	for i := 1; i < lim; i++ {
		base10 := strconv.FormatInt(int64(i), 10)
		if isPalindrome(base10) == false {
			continue
		}
		base2 := strconv.FormatInt(int64(i), 2)
		if isPalindrome(base2) == false {
			continue
		}
		// p(base10, " : ", base2)
		sum += i
	}
	return sum
}

func main() {

	p(findSumOfDualPalindromes(1000000))

}
