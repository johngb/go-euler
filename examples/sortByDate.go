package main

import "fmt"
import "time"
import "sort"

func main() {
	fmt.Println("Sorting an array of time")
	const shortForm = "Jan/02/2006"
	t1, _ := time.Parse(shortForm, "Feb/02/2014")
	t2, _ := time.Parse(shortForm, "Feb/02/1800")
	t3, _ := time.Parse(shortForm, "Feb/02/1999")
	t4, _ := time.Parse(shortForm, "Feb/02/2000")
	dates := []time.Time{t1, t2, t3, t4}
	for _, t := range dates {
		fmt.Println(t)
	}
	sort.Sort(ByDate(dates))
	fmt.Println("sorted:")
	for _, t := range dates {
		fmt.Println(t)
	}
}

type ByDate []time.Time

func (a ByDate) Len() int           { return len(a) }
func (a ByDate) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByDate) Less(i, j int) bool { return a[i].Before(a[j]) }
