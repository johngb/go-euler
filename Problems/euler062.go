/*
The cube, 41063625 (345^3), can be permuted to produce two other cubes:
56623104 (384^3) and 66430125 (405^3). In fact, 41063625 is the smallest cube
which has exactly three permutations of its digits which are also cube.

Find the smallest cube for which exactly five permutations of its digits are
cube.
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

func sortDigits(num int) int {

	numToRevSlice := func(n int) []int {
		ans := []int{}
		for n > 0 {
			ans = append(ans, n%10)
			n /= 10
		}
		return ans
	}

	popMinDigit := func(arr []int) (int, []int) {
		min := 99
		minIdx := 0
		for i := 0; i < len(arr); i++ {
			if arr[i] < min {
				min = arr[i]
				minIdx = i
			}
		}
		newArr := append(arr[0:minIdx], arr[minIdx+1:]...)
		return min, newArr
	}

	sortSlice := func(arr []int) int {
		num := 0
		digit := 0
		mul := 1
		for len(arr) > 0 {
			digit, arr = popMinDigit(arr)
			num += digit * mul
			mul *= 10
		}
		return num
	}

	digits := numToRevSlice(num)
	answer := (sortSlice(digits))
	return answer
}

func smallestCubeWithFiveCubicPermutations(lim int) int {
	defer timeTrack(time.Now(), "smallestCubeWithFiveCubicPermutations()")

	cubicPermMap := make(map[int][]int)
	for i := 1; i < lim; i++ {
		cube := i * i * i
		sortedDigits := sortDigits(cube)

		if cubicPermMap[sortedDigits] == nil {
			cubicPermMap[sortedDigits] = []int{1, cube}
		} else {
			cubicPermMap[sortedDigits][0] += 1
			cubicPermMap[sortedDigits] = append(cubicPermMap[sortedDigits], cube)
		}

		if cubicPermMap[sortedDigits][0] == 5 {
			// return the first number in the list, which should be the smallest
			return cubicPermMap[sortedDigits][1]
		}

	}
	p("Solution not found with lim =", lim)
	return 0
}

func main() {
	p(smallestCubeWithFiveCubicPermutations(10000))
}
