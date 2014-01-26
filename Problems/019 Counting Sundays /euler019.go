/*
You are given the following information, but you may prefer to do some research for yourself.

1 Jan 1900 was a Monday.
Thirty days has September,
April, June and November.
All the rest have thirty-one,
Saving February alone,
Which has twenty-eight, rain or shine.
And on leap years, twenty-nine.
A leap year occurs on any year evenly divisible by 4, but not on a century unless it is divisible by 400.
How many Sundays fell on the first of the month during the twentieth century (1 Jan 1901 to 31 Dec 2000)?
*/

package main

import (
	"fmt"
	"time"
)

var p = fmt.Println
var pf = fmt.Printf

// Used to time other functions to compare speed
func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s \n", name, elapsed)
}

func isLeapYear(year int) bool {
	// defer timeTrack(time.Now(), "isLeapYear()") // Timer function

	result := false

	switch {
	// definitely not a leap year
	case year%4 != 0:
		result = false
	// is a leap year if divisible by 400
	case year%400 == 0:
		result = true
	// not a leap year if divisible by 100
	case year%100 == 0:
		result = false
	// is a normal leap year
	case year%4 == 0:
		result = true
	default:
		p("Error with isLeapYear()")
		result = false
	}
	return result
}

func countMondayFirsts1() int {

	// 1 Jan 1900 was a monday
	dayCounter := 1
	mondayFirstCounter := 0

	for year := 1900; year < 2001; year++ {

		for month := 1; month <= 12; month++ {
			for day := 1; ; day++ {
				if day == 1 {
					// if the first of the month is a monday
					if year >= 1901 && dayCounter%7 == 1 {
						mondayFirstCounter++
					}
					dayCounter++
				} else if day == 29 {
					// if it's 29 February
					if month == 2 {
						if isLeapYear(year) {
							dayCounter++
							break
						} else {
							break
						}
					}
				} else if day == 30 {
					if month == 4 || month == 6 || month == 9 || month == 11 {
						dayCounter++
						break
					}
				} else if day == 31 {
					dayCounter++
					break
				} else {
					dayCounter++
				}
			}
		}
	}
	return mondayFirstCounter
}

// From spacepornGO on project Euler
func countMondayFirsts2() int {
	const dateFormat = "2006-Jan-02"
	date, err := time.Parse(dateFormat, "1901-Jan-01")
	if err != nil {
		fmt.Println(err)
	}
	endDate, err := time.Parse(dateFormat, "2000-Dec-31")
	if err != nil {
		fmt.Println(err)
	}
	// Loops while `date` is before `endDate`.
	count := 0
	for date.Before(endDate) {
		if date.Weekday() == time.Sunday {
			count++
		}
		// Adds 1 month.
		date = date.AddDate(0, 1, 0)
	}
	return count
}

func main() {
	p(isLeapYear(2100))
	p(countMondayFirsts1())
	p(countMondayFirsts2())
}
