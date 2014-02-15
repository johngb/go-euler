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

func makeLexicoPermList(arr []int) [][]int {

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

	answer := [][]int{}
	answer = append(answer, arr)
	notDone := false
	newArr := []int{}
	for {
		notDone, arr = nextLexicoPerm(arr)
		if notDone {
			newArr = make([]int, len(arr))
			copy(newArr, arr)
			answer = append(answer, newArr)

		} else {
			break
		}
		// p(answer)
	}
	return answer
}

func makeMapOfPolyNumberLists(arr []int) map[int][][]int {

	makePolyNumberList := func(polyNumber int) [][]int {
		const MIN = 1000
		const MAX = 10000
		list := [][]int{}
		num := 0
		for n := 1; ; n++ {
			switch polyNumber {
			case 3:
				num = n * (n + 1) / 2
			case 4:
				num = n * n
			case 5:
				num = n * (3*n - 1) / 2
			case 6:
				num = n * (2*n - 1)
			case 7:
				num = n * (5*n - 3) / 2
			case 8:
				num = n * (3*n - 2)
			default:
				p("type out of bounds.  Must be 3, 4, 5, 6, 7, or 8")
				return list
			}
			// if num is 4 digits
			if num < MAX && num >= MIN {

				list = append(list, []int{num})
			} else if num >= MAX {
				return list
			}
		}
		return list
	}

	answer := map[int][][]int{}

	for i := 0; i < len(arr); i++ {
		answer[arr[i]] = makePolyNumberList(arr[i])
	}
	return answer
}

// returns true if the last two digits of 'a' are the same as the first two digits of b
func isMatch(a, b int) bool {
	if a%100 == b/100 {
		return true
	}
	return false
}

func joinMatchedList(listA, listB [][]int) [][]int {
	solution := [][]int{}
	aIdx := len(listA[0]) - 1
	bIdx := 0
	for a := 0; a < len(listA); a++ {
		for b := 0; b < len(listB); b++ {
			if isMatch(listA[a][aIdx], listB[b][bIdx]) && listA[a][aIdx] != listB[b][bIdx] {
				newList := append(listA[a], listB[b]...)
				solution = append(solution, newList)
			}
		}
	}
	return solution
}

func matchBeginEnd(list [][]int) [][]int {
	answer := [][]int{}
	beginIdx := 0
	endIdx := len(list[0]) - 1
	for i := 0; i < len(list); i++ {
		if isMatch(list[i][endIdx], list[i][beginIdx]) {
			answer = append(answer, list[i])
		}

	}
	return answer
}

// 8128, 2882, 8281

// 8256, 5625, 2882, 8128, 2512, 1281
// 3		4	5		6	7		8
// 28684
// 1281, 8128, 2882, 8256, 5625, 2512

// order := 8, 6, 5, 3, 4, 7

func main() {
	arr := []int{3, 4, 5, 6, 7, 8}
	perms := makeLexicoPermList(arr)
	p("len perms:", len(perms))

	// perms := [][]int{[]int{8, 6, 5, 3, 4, 7}}
	polyMap := makeMapOfPolyNumberLists(arr)
	// p(polyMap)

	// answer := [][]int{}

	for i := 0; i < len(perms); i++ {
		// p("perms processing: ", perms[i])
		join1 := joinMatchedList(polyMap[perms[i][0]], polyMap[perms[i][1]])
		if len(join1) == 0 {
			continue
		}
		join2 := joinMatchedList(join1, polyMap[perms[i][2]])
		if len(join2) == 0 {
			continue
		}
		join3 := joinMatchedList(join2, polyMap[perms[i][3]])
		if len(join3) == 0 {
			continue
		}
		join4 := joinMatchedList(join3, polyMap[perms[i][4]])
		if len(join4) == 0 {
			continue
		}
		join5 := joinMatchedList(join4, polyMap[perms[i][5]])
		if len(join5) == 0 {
			continue
		}
		final := matchBeginEnd(join5)

		if len(final) != 0 {
			sum := 0
			for i := 0; i < len(final[0]); i++ {
				sum += final[0][i]
			}
			p(sum)
		}
	}
}
