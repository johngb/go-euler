/*
A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,

a2 + b2 = c2
For example, 32 + 42 = 9 + 16 = 25 = 52.

There exists exactly one Pythagorean triplet for which a + b + c = 1000.
Find the product abc.
*/

package main

import (
	"fmt"
	"math"
	"time"
)

var p = fmt.Println
var pf = fmt.Printf
var sqrt = math.Sqrt

// Used to time other functions to compare speed
func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s \n", name, elapsed)
}

func isPerfectSquare(n int) bool {
	if int(sqrt(float64(n)))*int(sqrt(float64(n))) == n {
		return true
	}
	return false
}

// Pure brute force
func PyTrip1(n int) int {
	// defer timeTrack(time.Now(), "PyTrip2()") // Timer function

	// maxC = n - minA - minB = n - 3
	for c := 1; c <= n; c++ {
		for b := 1; b <= n; b++ {
			for a := 1; a <= n; a++ {
				// if it's a perfect square and a + b + c = 1000
				if a+b+c == n {
					if a*a+b*b == c*c {
						// pf("a = %d, b = %d, c = %d\n", a, b, c)
						return a * b * c
					}
				}
			}
		}
	}
	return -1 // should never happen, so error
}

// Brute force with constraints to the values of a, b, and c
func PyTrip2(n int) int {
	// defer timeTrack(time.Now(), "PyTrip2()") // Timer function

	// From a + b + c == n, and given that a < b < c, the smallest value of c
	//  would be when a, b, and c are as close together as possible.
	//  so c-2 + c-1 + c = n.  Therefore c = n/3 + 1
	minC := (n / 3) + 1

	// maxC = n - minA - minB = n - 3
	for c := minC; c <= n-3; c++ {
		// From a^2 + b^2 = c^2, min b is when a = b-1.
		//  so, min b = sqrt((c-1)/2)
		minB := int(sqrt(float64((c - 1) / 2)))
		for b := minB; b <= c-1; b++ {
			for a := 1; a <= b-1; a++ {
				// if it's a perfect square and a + b + c = 1000
				if a+b+c == n {
					if a*a+b*b == c*c {
						// pf("a = %d, b = %d, c = %d\n", a, b, c)
						return a * b * c
					}
				}
			}
		}
	}
	return -1 // should never happen, so error
}

// Can be done more mathematically, by using Euclid's formula
// http://en.wikipedia.org/wiki/Pythagorean_triple#Generating_a_triple

func main() {
	p(PyTrip1(1000))
	p(PyTrip2(1000))
}
