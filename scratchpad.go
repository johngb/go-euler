package main

import (
	"fmt"
	"math"
	"time"
)

var p = fmt.Println
var pf = fmt.Printf

func proc1() {
	p("process 1 started")
	x := math.Sqrt(9865651)
	p(x)
	p("process 1 done")
}

func proc2() {
	p("process 2 running")
}

func main() {

	// go proc1()
	// proc2()

	timerChan := make(chan time.Time)
	go func() {
		time.Sleep(3000)
		timerChan <- time.Now() // send time on timerChan
	}()
	// Do something else; when ready, receive.
	// Receive will block until timerChan delivers.
	// Value sent is other goroutine's completion time.
	completedAt := <-timerChan
	p(completedAt)

}
