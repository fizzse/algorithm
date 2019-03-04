package main

import "fmt"

func BinarySearch(arr []int, dest int) int {
	low := 0
	high := len(arr)

	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == dest {
			return mid
		} else if arr[mid] > dest {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 13, 15, 17, 19, 31, 34, 56, 78, 89, 100}
	dest := 31

	index := BinarySearch(arr, dest)
	if index < 0 {
		fmt.Println("no data in this set")
	}

	fmt.Printf("data:[%d] find in this set index:[%d]\n", dest, index)
}
