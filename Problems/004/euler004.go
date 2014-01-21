package main

import (
	"fmt"
	"time"
)

var pl = fmt.Println
var pf = fmt.Printf

// Used to time other functions to compare speed
func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s \n", name, elapsed)
}

func isPalindrome(num int) bool {
	// if num%10 == 0, can't be a palindrome as number can't start with 0
	if num%10 == 0 {
		return false
	}
	// split num into a slice with each element 0-9 representing a digit
	// this creates a reversed slice, but if it's a plaindrome, it won't matter
	slice := make([]int, 0)
	for n := num; n > 0; {
		slice = append(slice, n%10)
		n = n / 10
	}
	// slice.first index = 0, slice.last index = length slice
	first := 0
	last := len(slice) - 1 // -1 to account for index starting at 0
	// while slice.first.index <= slice.last index
	for first <= last {
		// if slice.first != slice.last, not a palindrome
		if slice[first] != slice[last] {
			return false
		}
		// increment slice.first index
		first++
		// decrement slice.last index
		last--
	}
	// is a palindrome
	return true
}

func palindrome1() [3]int {
	// defer timeTrack(time.Now(), "palindrome1()") // Timer function
	var prod int
	// set up an empty array to hold the answer
	array := [3]int{0, 0, 0}
	// loop through first 3 digit number
	for n1 := 100; n1 <= 999; n1++ {
		// loop through second 3 digit number
		for n2 := 100; n2 <= 999; n2++ {
			prod = n1 * n2
			// if product is an palindrome
			if isPalindrome(prod) {
				// if palindrome is larger than the one in array, replace it
				if prod > array[0] {
					array[0] = prod
					array[1] = n1
					array[2] = n2
				}
			}
		}
	}
	// pl(array)
	return array
}

func palindrome2() [3]int {
	// defer timeTrack(time.Now(), "palindrome2()") // Timer function
	var prod int
	// set up an empty array to hold the answer
	array := [3]int{0, 0, 0}
	maxPalindrome := 0
	// loop through first 3 digit number
	for n1 := 100; n1 <= 999; n1++ {
		// loop through second 3 digit number
		for n2 := 100; n2 <= 999; n2++ {
			prod = n1 * n2
			// if product is an palindrome
			if prod > maxPalindrome && isPalindrome(prod) {
				// if palindrome is larger than the one in array, replace it
				if prod > array[0] {
					array[0] = prod
					maxPalindrome = prod
					array[1] = n1
					array[2] = n2
				}
			}
		}
	}
	// pl(array)
	return array
}

func palindrome3() [3]int {
	// defer timeTrack(time.Now(), "palindrome3()") // Timer function

	var prod int
	// set up an empty array to hold the answer
	array := [3]int{0, 0, 0}
	maxPalindrome := 0
	// loop through first 3 digit number
	for n1 := 999; n1 >= 100; n1-- {
		// loop through second 3 digit number
		for n2 := 999; n2 >= 100; n2-- {
			prod = n1 * n2
			// if product is an palindrome
			if prod > maxPalindrome && isPalindrome(prod) {
				// if palindrome is larger than the one in array, replace it
				if prod > array[0] {
					array[0] = prod
					maxPalindrome = prod
					array[1] = n1
					array[2] = n2
				}
			}
		}
	}
	// pl(array)
	return array
}

func main() {
	// pl(isPalindrome(12345678987654321)) // call the function being tested
	palindrome1()
	palindrome2()
	palindrome3()
}
