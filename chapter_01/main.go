package main

import "fmt"

// Example Binary search
func binarySearch(arr []int, item int) int {
	low := 0
	hight := len(arr) - 1
	for low <= hight {
		mid := (low + hight) / 2
		guess := arr[mid]

		if guess == item {
			return guess
		}

		if guess < item {
			low = mid + 1
		} else {
			hight = mid - 1
		}
	}

	return -1
}

func main() {
	// Binary search
	arr := [10]int{12, 22, 33, 44, 55, 66, 77, 88, 99, 100}
	var item int
	fmt.Print("Example Binary Search Array [12, 22, 33, 44, 55, 66, 77, 88, 99, 100]\nEnter number: ")
	_, err := fmt.Scan(&item)
	if err != nil {
		fmt.Println("Something is wrong!")
	}
	fmt.Printf("Result binary Searching for item %d is: %d", item, binarySearch(arr[:], item))

}
