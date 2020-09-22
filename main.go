package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var n int
	fmt.Println("バブルソート")
	n = 1
	for i := 0; i < 6; i++ {
		a := ramdom(n, n)
		fmt.Println("長さ:", len(a))
		BubbleSort(a)
		n = n * 10
	}
	/*
		fmt.Println("選択ソート")
		n = 1
		for i := 0; i < 6; i++ {
			a := ramdom(n, n)
			SelectionSort(a)
			n = n * 10
		}
		fmt.Println("挿入ソート")
		n = 1
		for i := 0; i < 6; i++ {
			a := ramdom(n, n)
			InsertSort(a)
			n = n * 10
		}
		fmt.Println("シェルソート")
		n = 1
		for i := 0; i < 6; i++ {
			a := ramdom(n, n)
			ShellSort(a)
			n = n * 10
		}
		fmt.Println("マージソート")
		n = 10
		for i := 0; i < 6; i++ {
			a := ramdom(n, n)
			start := time.Now()
			MergeSort(a)
			end := time.Now()
			fmt.Printf("%f秒\n", (end.Sub(start)).Seconds())
			n = n * 10
		}
		fmt.Println("クイックソート")
		n = 1
		for i := 0; i < 6; i++ {
			a := ramdom(n, n)
			start := time.Now()
			QuickSort(a)
			end := time.Now()
			fmt.Printf("%f秒\n", (end.Sub(start)).Seconds())
			n = n * 10
		}
		fmt.Println("ヒープソート")
		n = 1
		for i := 0; i < 6; i++ {
			a := ramdom(n, n)
			start := time.Now()
			HeapSort(a)
			end := time.Now()
			fmt.Printf("%f秒\n", (end.Sub(start)).Seconds())
			n = n * 10
		}
	*/
}

func ramdom(a, b int) []int {
	num := []int{}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < a; i++ {
		num = append(num, rand.Intn(b))
	}
	return num
}
