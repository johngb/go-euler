package main

import (
	. "fmt"
	. "math"
	. "strconv"
)

func main() {
	total := 0
	number := FormatFloat(Pow(2, 1000), 'f', 0, 64)

	for _, v := range number {
		digit, _ := Atoi(string(v))
		total += digit
	}
	Println(total)
}
