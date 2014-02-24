/*
Using the numbers 1 to 10, and depending on arrangements, it is possible to
form 16- and 17-digit strings. What is the maximum 16-digit string for a
"magic" 5-gon ring?
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

func makeLexicoPermListOf5Gongs() [][]int {

	// nextLexicoPerm returns true if there is another lexicographic permutation
	// possible, as well as a slice showing the permutation
	nextLexicoPerm := func(inputArr []int) (bool, []int) {
		// defer timeTrack(time.Now(), "nextLexicoPerm()") // Timer function

		length := len(inputArr)
		arr := make([]int, length)
		copy(arr, inputArr)
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

	is5Gong := func(a []int) bool {
		sum := a[5] + a[0] + a[1]
		if a[6]+a[1]+a[2] != sum {
			return false
		} else if a[7]+a[2]+a[3] != sum {
			return false
		} else if a[8]+a[3]+a[4] != sum {
			return false
		} else if a[9]+a[4]+a[0] != sum {
			return false
		}
		return true
	}

	make5GongString := func(a []int) []int {

		rows := [][]int{
			[]int{a[5], a[0], a[1]},
			[]int{a[6], a[1], a[2]},
			[]int{a[7], a[2], a[3]},
			[]int{a[8], a[3], a[4]},
			[]int{a[9], a[4], a[0]},
		}
		// find the row with the lowest outer digit
		min := 99
		minIdx := 0
		for k, v := range rows {
			if v[0] < min {
				min = v[0]
				minIdx = k
			}
		}
		result := []int{}
		for i := 0; i < len(rows); i++ {
			result = append(result, rows[(minIdx+i)%len(rows)]...)
		}
		return result
	}

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	answer := [][]int{}
	notDone := false

	for {
		notDone, arr = nextLexicoPerm(arr)
		if notDone {
			if is5Gong(arr) {
				// JGB make 5gong string
				newArr := make5GongString(arr)
				answer = append(answer, newArr)
			}
		} else {
			break
		}
		// p(answer)
	}
	return answer
}

func listToString(list []int) string {
	str := ""
	for i := 0; i < len(list); i++ {
		str += strconv.Itoa(list[i])
	}
	return str
}

func euler068() int {
	defer timeTrack(time.Now(), "euler068()")

	perms := makeLexicoPermListOf5Gongs()
	maxNum := 0

	for _, v := range perms {
		str := listToString(v)
		// p("str = ", str)
		if len(str) == 16 {
			// p(str)
			num, _ := strconv.Atoi(str)
			// p(num)
			if num > maxNum {
				maxNum = num
			}
		}
	}

	return maxNum
}

func main() {

	p(euler068())

}
