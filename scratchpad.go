package main

// import "fmt"

const n = 10000

func withCap() {
	b := make([]int, 0, n*10)
	for i := 0; i < n; i++ {
		b = append(b, i)
	}
}

func noCap() {
	b := make([]int, 0)
	for i := 0; i < n; i++ {
		b = append(b, i)
		// fmt.Println("len:", len(b), "cap:", cap(b))
	}
}

func main() {
	noCap()
}
