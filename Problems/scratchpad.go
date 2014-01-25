package main

import (
	"fmt"
)

var p = fmt.Println

var arr [][]uint64

const n = 2

func main() {
	arr = make([][]uint64, n+1)
	for i := range arr {
		arr[i] = make([]uint64, n+1)
		for j := range arr[i] {
			arr[i][j] = uint64(0)
		}
	}
	arr[n][n] = uint64(1)
	println(noOfRoutes(0, 0))

	for i := range arr {
		p(arr[i])
	}
}

func noOfRoutes(i, j int) uint64 {
	if arr[i][j] != uint64(0) {
		return arr[i][j]
	}
	var result uint64
	if i < n && j < n {
		result = noOfRoutes(i+1, j) + noOfRoutes(i, j+1)
	} else if i < n && j == n {
		result = noOfRoutes(i+1, j)
	} else {
		result = noOfRoutes(i, j+1)
	}
	arr[i][j] = result
	return result
}
