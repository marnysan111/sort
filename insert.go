package main

import (
	"fmt"
	"time"
)

func InsertSort(a []int) []int {
	start := time.Now()
	for i := 1; i < len(a); i++ {
		for j := 0; j < i; j++ {
			if a[i-j-1] > a[i-j] {
				a[i-j-1], a[i-j] = a[i-j], a[i-j-1]
			} else {
				break
			}
		}
	}
	end := time.Now()
	fmt.Printf("%f秒\n", (end.Sub(start)).Seconds())
	return a
}
