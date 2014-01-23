package main

import (
	"fmt"
	"strconv"
	// "strings"
	"time"
)

var p = fmt.Println
var pf = fmt.Printf

// Used to time other functions to compare speed
func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s \n", name, elapsed)
}

// convert a char to an int.  No checking int place
func ord(c byte) int {
	return int(c - '0')
}

func largestProduct1(N string) int {
	// defer timeTrack(time.Now(), "largestProduct1()") // Timer function

	product, largestProduct := 0, 0
	char1, char2, char3, char4, char5 := 0, 0, 0, 0, 0
	runs := 0

	for idx := 0; idx < len(N)-5; idx++ {
		char1, _ = strconv.Atoi(N[idx+0 : idx+1])
		char2, _ = strconv.Atoi(N[idx+1 : idx+2])
		char3, _ = strconv.Atoi(N[idx+2 : idx+3])
		char4, _ = strconv.Atoi(N[idx+3 : idx+4])
		char5, _ = strconv.Atoi(N[idx+4 : idx+5])
		product = char1 * char2 * char3 * char4 * char5
		// pf("%d + %d + %d + %d + %d = %d\n", char1, char2, char3, char4, char5, product)

		if product > largestProduct {
			largestProduct = product
		}
		runs++
	}
	// p(runs)
	return largestProduct
}

func largestProduct2(N string) int {
	// defer timeTrack(time.Now(), "largestProduct2()") // Timer function

	product, largestProduct := 0, 0
	runs := 0

	for idx := 0; idx < len(N)-5; idx++ {
		product = ord(N[idx]) * ord(N[idx+1]) * ord(N[idx+2]) * ord(N[idx+3]) * ord(N[idx+4])
		// p(product)

		if product > largestProduct {
			largestProduct = product
		}
		runs++
	}
	// p(runs)
	return largestProduct
}

func largestProduct3(N string) int {
	// defer timeTrack(time.Now(), "largestProduct3()") // Timer function

	product, largestProduct := 0, 0
	char1 := 0
	// runs := 0

	for idx := 0; idx < len(N)-5; idx++ {
		char1 = ord(N[idx])
		if char1 == 0 {
			// if a char is 0, the next 5 loops will be 0, so skip this and the next 4
			idx += 4
			continue
		}
		product = char1 * ord(N[idx+1]) * ord(N[idx+2]) * ord(N[idx+3]) * ord(N[idx+4])

		if product > largestProduct {
			largestProduct = product
		}
		// runs++
	}
	// p(runs)
	return largestProduct
}

func main() {

	const N = "73167176531330624919225119674426574742355349194934" +
		"96983520312774506326239578318016984801869478851843" +
		"85861560789112949495459501737958331952853208805511" +
		"12540698747158523863050715693290963295227443043557" +
		"66896648950445244523161731856403098711121722383113" +
		"62229893423380308135336276614282806444486645238749" +
		"30358907296290491560440772390713810515859307960866" +
		"70172427121883998797908792274921901699720888093776" +
		"65727333001053367881220235421809751254540594752243" +
		"52584907711670556013604839586446706324415722155397" +
		"53697817977846174064955149290862569321978468622482" +
		"83972241375657056057490261407972968652414535100474" +
		"82166370484403199890008895243450658541227588666881" +
		"16427171479924442928230863465674813919123162824586" +
		"17866458359124566529476545682848912883142607690042" +
		"24219022671055626321111109370544217506941658960408" +
		"07198403850962455444362981230987879927244284909188" +
		"84580156166097919133875499200524063689912560717606" +
		"05886116467109405077541002256983155200055935729725" +
		"71636269561882670428252483600823257530420752963450"

	p(largestProduct1(N))
	p(largestProduct2(N))
	p(largestProduct3(N))

}
