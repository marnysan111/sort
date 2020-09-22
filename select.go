package main

import (
	"fmt"
	"time"
)

func Min(a []int) (idx, n int) {
	n = a[0]
	idx = 0
	for i, v := range a {
		if n > v {
			n = v
			idx = i
		}
	}
	return
}

func SelectionSort(a []int) []int {
	start := time.Now()
	for i, _ := range a {
		idx, _ := Min(a[i:len(a)])
		a[i], a[i+idx] = a[i+idx], a[i]
	}
	end := time.Now()
	fmt.Printf("%f秒\n", (end.Sub(start)).Seconds())
	return a
}
