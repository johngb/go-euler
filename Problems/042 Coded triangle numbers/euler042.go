/*
The nth term of the sequence of triangle numbers is given by, tn = ½n(n+1); so
the first ten triangle numbers are:

1, 3, 6, 10, 15, 21, 28, 36, 45, 55, ...

By converting each letter in a word to a number corresponding to its
alphabetical position and adding these values we form a word value. For
example, the word value for SKY is 19 + 11 + 25 = 55 = t10. If the word value
is a triangle number then we shall call the word a triangle word.

Using words.txt (right click and 'Save Link/Target As...'), a 16K text file
containing nearly two-thousand common English words, how many are triangle
words?
*/

package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

var p = fmt.Println
var pf = fmt.Printf

// timeTrack is used for basic benchmarking in other functions
func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s \n", name, elapsed)
}

// nameLetterSun returns the sum of the numerical values of the letters in a
// string.  Using A = 1, B = 2, etc.
func nameLetterSum(name string) int {
	sum := 0
	for i := 0; i < len(name); i++ {
		sum += int(name[i] - 'A' + 1)
	}
	return sum
}

// makeTriangleNumberMap returns a map of triangle numbers up to the 'maxN'th number
func makeTriangleNumberMap(maxN int) map[int]bool {

	triangleNumberMap := make(map[int]bool)

	for n := 1; n <= maxN; n++ {
		triangleNumber := n * (n + 1) / 2
		triangleNumberMap[triangleNumber] = true
	}

	return triangleNumberMap
}

func readTextFileAndProcess(fileName string) []string {
	defer timeTrack(time.Now(), "readTextFileAndProcess()") // Timer function

	fileBuf, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	// remove leading and trailing Unicode code points
	fileStr := strings.Trim(string(fileBuf), "")
	// remove all quotation marks (")
	fileStr = strings.Replace(fileStr, "\"", "", -1)
	// remove all newlines
	fileStr = strings.Replace(fileStr, "\n", "", -1)
	nameList := strings.Split(fileStr, ",")

	return nameList
}

// countTriangleWords returns the number of words in the file 'fileName' whose
// sum of letter position values is a triangle number
func countTriangleWords(fileName string) int {
	defer timeTrack(time.Now(), "countTriangleWords()")

	nameList := readTextFileAndProcess(fileName)
	triangleMap := makeTriangleNumberMap(50)

	triangleWordCount := 0

	for i := 0; i < len(nameList); i++ {
		letterSum := nameLetterSum(nameList[i])
		if triangleMap[letterSum] == true {
			triangleWordCount++
		}
	}
	return triangleWordCount
}

func main() {

	// p(readTextFileAndProcess("words.txt"))
	// p(nameLetterSum("SKY"))
	// p(makeTriangleNumberMap(10))
	p(countTriangleWords("words.txt"))

}
