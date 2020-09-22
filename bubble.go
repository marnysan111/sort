package main

import "fmt"

func BubbleSort(a []int) {
	//start := time.Now()
	var count int
	for i := 0; i < len(a)-1; i++ {
		for j := 0; j < len(a)-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
			count++
		}
	}
	fmt.Println(count)
	//end := time.Now()
	//output := end.Sub(start).Seconds()
	//fmt.Printf("%f秒\n", output)
}
