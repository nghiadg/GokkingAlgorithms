package main

import "fmt"

// Selection sort
func findSmallest(arr []int) int {
	smallest := arr[0]
	smallestIndex := 0

	for i, v := range arr {
		if v < smallest {
			smallest = v
			smallestIndex = i
		}
	}

	return smallestIndex
}

func selectionSort(arr []int) []int {
	var newArr []int
	for range len(arr) {
		smallest := findSmallest(arr)
		newArr = append(newArr, arr[smallest])
		// remove item from arr
		arr = append(arr[:smallest], arr[smallest+1:]...)

	}

	return newArr
}

func main() {
	fmt.Println("Example Selection Sort:")
	arr := [10]int{152, 34, 373, 67, 34, 88, 324, 86, 32, 56}
	fmt.Println("Sorted Array: ", selectionSort(arr[:]))
}
