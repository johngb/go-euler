/*
The prime factors of 13195 are 5, 7, 13 and 29.
What is the largest prime factor of the number 600851475143 ?
*/

package main

// package euler003

import (
	"fmt"
	"time"
)

// Used to time other functions to compare speed
func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s \n", name, elapsed)
}

// Inefficient function to check if a number is prime
// Could be improved by not checking multiples of i
func isPrime(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func eratosthenes1(num int) []int {
	// initialise an array of with num elements
	sieve := make([]bool, num+2)
	// index = 2
	var index int = 2

	// while index <= num
	for index <= num {

		// make each element <= num that is a multiple of index = 1
		var i int = 1
		for i = 2; i*index <= num; i++ {
			sieve[i*index] = true
		}

		// increment index to next element of the array that = 0
		for {
			index += 1
			if sieve[index] == false {
				break
			}
		}
	}

	// if an element of the array == 0, add its index to a list
	primeFactors := make([]int, 0)
	for x := 2; x <= num; x++ {
		if sieve[x] == false {
			primeFactors = append(primeFactors, x)
		}
	}
	// return the list
	return primeFactors
}

func eratosthenes2(num int) []int { // initialise an array of with num elements sieve := make([]bool, num+2)
	var sieve [1000]bool // be sure that 100 >= n
	// index = 2
	var index int = 2

	// intialise slice to contain prime numbers
	primes := make([]int, 0)

	// while index <= num
	for index <= num {
		// index should be a prime number, so add to list of primes
		primes = append(primes, index)

		// make each element <= num that is a multiple of index = 1
		var i int = 1
		for i = 2; i*index <= num; i++ {
			sieve[i*index] = true
		}

		// increment index to next element of the array that = 0
		for {
			index += 1
			if sieve[index] == false {
				break
			}
		}
	}

	// return the list of primes
	return primes
}

func eratosthenes3(num int) []int {
	// initialise an array of with num elements
	sieve := make([]bool, num+2)
	// intialise slice to contain prime numbers
	primes := make([]int, 0)

	// index = 2
	// while i <= num
	// for i := 2; i <= int(math.Sqrt(float64(num)))+1; i++ {
	for i := 2; i <= num; i++ {

		if sieve[i] == false {

			// i should be a prime number, so add to list of primes
			primes = append(primes, i)

			for j := 2; i*j <= num; j++ {
				sieve[i*j] = true
			}
		}
	}

	// return the list of primes
	return primes
}

// Works but is far too slow for large numbers
func biggestPrimeFactor1(num int) int {
	for i := num; i >= 1; i-- {
		// if it is a factor AND it is a prime
		if num%i == 0 && isPrime(i) == true {
			return i
		}
	}
	return 0
}

func biggestPrimeFactor2(num int) int {
	n := num
	d := 2
	for {
		if n%d != 0 {
			d++
		} else {
			// fmt.Printf("n = %d, d = %d, ", n, d)
			n = n / d
			// fmt.Printf("new n = %d\n", n)
		}
		if d > n {
			break
		}
	}
	return d
}

// Best solution of all
func biggestPrimeFactor3(num int) int {
	n := num
	d := 2
	for {
		if n%d == 0 {
			n = n / d
		}
		// There can't be another factor greater than d if d > n, so stop
		if d > n {
			break
		}
		d = d + 1
	}
	return d
}

func smallestFactor1(num int) int {
	// defer timeTrack(time.Now(), "smallestFactor1") // Timer function

	for i := 2; i <= num; i++ {
		// if it is a factor AND it is a prime
		if num%i == 0 {
			return i
		}
	}
	return 0 // should never happen
}

func factors1(num int) []int {
	// defer timeTrack(time.Now(), "factors1()") // Timer function

	n := num
	factors := make([]int, 0)
	for {
		f := smallestFactor1(n)
		factors = append(factors, f)

		// new number to find the factors of
		n = n / f
		// no more factors
		if n == 1 {
			return factors
		}
	}
}

// Method using recursion with a global.
// globals are BAD.
// Keeping for a reference
var factorSlice2 []int

func factors2(num int) []int {
	// defer timeTrack(time.Now(), "factors2()") // Timer function

	// var factorSlice2 []int
	for i := 2; i <= num; i++ {

		// if i is a factor of num
		if num%i == 0 {
			// add i to factorSlice2
			factorSlice2 = append(factorSlice2, i)

			// new number to factorise
			temp := num / i

			// get the next smallest factor
			if temp == num {
				return factorSlice2
			}
			return factors2(temp)
		}
	}
	return factorSlice2
}

// method using internal recursion
func factors3(num int) []int {
	// defer timeTrack(time.Now(), "factors3()") // Timer function

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
				factorSlice = append(factorSlice, factors3(temp)...)
			}
			// factorSlice will now contain all factors of whatever it was called with
			// so return the solution
			return factorSlice
		}
	}
	// should never need this return
	return []int{0}
}

// Optimisation of factors1()
func factors4(num int) []int {
	// defer timeTrack(time.Now(), "factors4()") // Timer function

	factors := make([]int, 0)
	n := num
	d := 2
	for n > 1 {
		if n%d == 0 {
			factors = append(factors, d)
			n = n / d
		}
		// There can't be another factor greater than d if d > n, so stop
		if d > n {
			// fmt.Printf("d * d = %d, num = %d\n", d*d, num)
			break
		}
		d = d + 1
	}
	return factors
}

func main() {
	const n = 100
	// fmt.Println(eratosthenes1(n))
	// fmt.Println(eratosthenes2(n))
	// fmt.Println(eratosthenes3(n))
	fmt.Println(eratosthenes1(n))
	fmt.Println(eratosthenes2(n))
	fmt.Println(eratosthenes3(n))

	// const p = 600851475143 //the value to find the biggest prime factor of
	// const p = 266898757345
	// const p = 3559 * 3571
	// const p = 156903 // a prime * 3
	// const p = 12876876

	// fmt.Printf("Is '%d' prime: %t\n", p, isPrime(p))
	// fmt.Printf("Biggest prime factor in %d is: %d\n", p, biggestPrimeFactor1(p))
	// fmt.Printf("Smallest factor in %d is: %d\n", p, smallestFactor1(p))
	// fmt.Println(factors3(p))
	// fmt.Println(factors4(p))
	// fmt.Println(biggestPrimeFactor2(p))
	// fmt.Println(biggestPrimeFactor3(p))
}
