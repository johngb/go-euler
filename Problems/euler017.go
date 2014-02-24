/*If the numbers 1 to 5 are written out in words: one, two, three, four, five,
then there are 3 + 3 + 5 + 4 + 4 = 19 letters used in total.

If all the numbers from 1 to 1000 (one thousand) inclusive were written out in
words, how many letters would be used?

NOTE: Do not count spaces or hyphens. For example, 342 (three hundred and
forty-two) contains 23 letters and 115 (one hundred and fifteen) contains 20
letters. The use of "and" when writing out numbers is in compliance with
British usage.*/

package main

import (
	"fmt"
	"strings"
	"time"
)

var p = fmt.Println
var pf = fmt.Printf

// Used to time other functions to compare speed
func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s \n", name, elapsed)
}

func numberToWords(i int) string {

	wordMap := map[int]string{
		0:    "",
		2:    "two",
		1:    "one",
		3:    "three",
		4:    "four",
		5:    "five",
		6:    "six",
		7:    "seven",
		8:    "eight",
		9:    "nine",
		10:   "ten",
		11:   "eleven",
		12:   "twelve",
		13:   "thirteen",
		14:   "fourteen",
		15:   "fifteen",
		16:   "sixteen",
		17:   "seventeen",
		18:   "eighteen",
		19:   "nineteen",
		20:   "twenty",
		30:   "thirty",
		40:   "forty",
		50:   "fifty",
		60:   "sixty",
		70:   "seventy",
		80:   "eighty",
		90:   "ninety",
		100:  "hundred",
		1000: "thousand",
	}

	word := ""
	units := 0
	tens := 0
	hundreds := 0

	switch {

	// 1 -> 19
	case i < 20:
		word += wordMap[i]

	// 20 -> 99
	case i < 100:
		units = i % 10
		tens = (i / 10) * 10
		word += wordMap[tens] + " " + wordMap[units]

	// 100 - 999
	case i < 1000:
		hundreds = (i / 100) * 100
		leftover := i - hundreds

		// exactly hundreds, so no "and" needed
		if leftover == 0 {
			word += wordMap[hundreds/100] + " " + wordMap[100]
			// need to include an "and" as well as the rest of the words
		} else {
			word += wordMap[hundreds/100] + " " + wordMap[100] + " and " + numberToWords(leftover)
		}

	// 1000
	case i == 1000:
		word += wordMap[1] + " " + wordMap[1000]

	// error
	default:
		word = "error: number out of range"
	}

	return word
}

func countLetters(lim int) int {
	defer timeTrack(time.Now(), "countLetters()")

	count := 0
	word := ""
	wordlength := 0
	for i := 1; i <= lim; i++ {
		word = numberToWords(i)
		wordlength = len(strings.Replace(word, " ", "", -1))
		// pf("%d: length = %d\n", i, wordlength)
		count += wordlength
	}
	return count
}

func main() {

	p(countLetters(1000))
}
