/*
The first two consecutive numbers to have two distinct prime factors are:

14 = 2 × 7
15 = 3 × 5

The first three consecutive numbers to have three distinct prime factors are:

644 = 2² × 7 × 23
645 = 3 × 5 × 43
646 = 2 × 17 × 19.

Find the first four consecutive integers to have four distinct prime factors.
What is the first of these numbers?
*/

// JGB: Lots of room for improvement here.  Mostly in finding the factors.

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

// method using internal recursion.  Room for optimisation here
func primeFactors(num int) []int {
	// defer timeTrack(time.Now(), "primeFactors()") // Timer function

	var factorSlice []int
	for i := 2; i <= num; i++ {

		// if i is a factor of num
		if num%i == 0 {
			// add i to factorSlice
			factorSlice = append(factorSlice, i)

			// new number to factorise
			temp := num / i

			if temp > 1 {
				// get the rest of the factors via recursion and append to list
				factorSlice = append(factorSlice, primeFactors(temp)...)
			}
			// factorSlice will now contain all factors of whatever it was called with
			// so return the solution
			return factorSlice
		}
	}
	return []int{0}
}

func uniqueFactors(num int) []int {
	// defer timeTrack(time.Now(), "uniqueFactors()")

	primeFactorList := primeFactors(num)

	// make the map of factors
	factorMap := make(map[int]int)
	for i := 0; i < len(primeFactorList); i++ {
		factorMap[primeFactorList[i]]++
	}

	//expand duplicate factors
	uniqueFactorList := []int{}
	for k, v := range factorMap {
		uniqueFactorList = append(uniqueFactorList, int(math.Pow(float64(k), float64(v))))
	}
	return uniqueFactorList
}

func solve047(lim int) int {
	defer timeTrack(time.Now(), "solve047()") // Timer function

	factorMap := make(map[int]bool)
	currentFactors := []int{}
	const NUM = 4
	count := 0

	reset := func() {
		factorMap = make(map[int]bool)
		count = 0
	}
	loadCurrentToMap := func() bool {
		for x := 0; x < len(currentFactors); x++ {
			if factorMap[currentFactors[x]] == true {
				return false
			} else {
				factorMap[currentFactors[x]] = true
			}
		}
		return true
	}

	for i := 10; i < lim; i++ {
		currentFactors = uniqueFactors(i)
		if len(currentFactors) != NUM {
			reset()
			continue
		}

		if loadCurrentToMap() == false {
			reset()
			continue
		}

		count++
		if count == NUM {
			return i - NUM + 1
		}
	}
	return 0
}

func main() {

	// p(primeFactors(28622480))
	// p(uniqueFactors(286229))
	p(solve047(150000))

}
