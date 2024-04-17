package main

import (
	"fmt"
)

func quicksort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	less := []int{}
	greater := []int{}

	pivot := arr[0]

	for _, v := range arr[1:] {
		if v < pivot {
			less = append(less, v)
		} else {
			greater = append(greater, v)
		}
	}

	return append((append(quicksort(less), []int{pivot}...)), quicksort(greater)...)
}

func main() {
	arr := [10]int{1, 3, 7, 2, 6, 9, 3, 7, 4, 6}

	fmt.Println("Quicksort result: ", quicksort(arr[:]))
}
