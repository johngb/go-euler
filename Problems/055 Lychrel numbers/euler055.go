/*
If we take 47, reverse and add, 47 + 74 = 121, which is palindromic.

Not all numbers produce palindromes so quickly. For example,

349 + 943 = 1292,
1292 + 2921 = 4213
4213 + 3124 = 7337

That is, 349 took three iterations to arrive at a palindrome.

Although no one has proved it yet, it is thought that some numbers, like 196,
never produce a palindrome. A number that never forms a palindrome through the
reverse and add process is called a Lychrel number. Due to the theoretical
nature of these numbers, and for the purpose of this problem, we shall assume
that a number is Lychrel until proven otherwise. In addition you are given
that for every number below ten-thousand, it will either (i) become a
palindrome in less than fifty iterations, or, (ii) no one, with all the
computing power that exists, has managed so far to map it to a palindrome. In
fact, 10677 is the first number to be shown to require over fifty iterations
before producing a palindrome: 4668731596684224866951378664 (53 iterations,
28-digits).

Surprisingly, there are palindromic numbers that are themselves Lychrel
numbers; the first example is 4994.

How many Lychrel numbers are there below ten-thousand?
*/

package main

import (
	"fmt"
	"math/big"
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

// reverseString returns a new string which is a reverse of the input 'str'
func reverseString(str string) string {
	newStr := ""

	for idx := len(str); idx > 0; idx-- {
		newStr += str[idx-1 : idx]
	}
	return newStr
}

func nextLychrelIteration(str string) (bool, string) {
	strInt := big.NewInt(0)
	strInt.SetString(str, 10)

	revStr := reverseString(str)
	revStrInt := big.NewInt(0)
	revStrInt.SetString(revStr, 10)

	result := big.NewInt(0)
	result.Add(strInt, revStrInt)

	resultStr := result.String()
	// p("resultStr:", resultStr)
	if isPalindrome(resultStr) {
		return true, resultStr
	}
	return false, resultStr
}

func findLychrelBelow(lim int) int {
	defer timeTrack(time.Now(), "findLychrelBelow()")

	lychrelCounter := 0

	for i := 1; i < lim; i++ {
		potLychrel := strconv.Itoa(i)
		notLychrel := false
		// run at most 50 times
		for iteration := 1; iteration <= 50; iteration++ {
			notLychrel, potLychrel = nextLychrelIteration(potLychrel)
			if notLychrel {
				break
			}

		}
		if notLychrel == false {
			lychrelCounter++
		}
	}
	return lychrelCounter

}

func main() {

	p(findLychrelBelow(10000))
}
