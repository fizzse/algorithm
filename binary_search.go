package main

import "fmt"

// 普通二分查找
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

// 在有序 重复数组中查找第一个等于指定值的数
func FirstEqual(arr []int, dest int) int {
	high := len(arr) - 1
	low := 0

	for low <= high {
		mid := low + (high-low)/2
		if dest > arr[mid] {
			low = mid + 1
		} else if dest < arr[mid] {
			high = mid - 1
		} else {
			if mid == 0 || arr[mid-1] != dest {
				return mid
			} else {
				high = mid - 1
			}
		}
	}

	return -1
}

// 在有序 重复数组中查找最后一个等于指定值的数
func LastEqual(arr []int, dest int) int {
	high := len(arr) - 1
	low := 0

	for low <= high {
		mid := low + (high-low)/2
		if dest > arr[mid] {
			low = mid + 1
		} else if dest > arr[mid] {
			high = mid - 1
		} else {
			if mid == len(arr)-1 || arr[mid+1] != dest {
				return mid
			} else {
				low = mid + 1
			}
		}
	}

	return -1
}

// 在有序 重复数组中查找第一个大于等于指定值的数
func FirstMore(arr []int, dest int) int {
	high := len(arr) - 1
	low := 0

	for low <= high {
		mid := low + (high-low)/2
		if dest <= arr[mid] {
			if mid == 0 || arr[mid-1] < dest {
				return mid
			} else {
				high = mid - 1
			}

		} else if dest > arr[mid] {
			low = mid + 1
		}
	}

	return -1
}

// TODO. 在有序 重复数组中查找最后一个小于等于指定值的数
func LastLess(arr []int, dest int) int {
	high := len(arr) - 1
	low := 0

	for low <= high {
		mid := low + (high-low)/2
		fmt.Println(mid)
	}

	return -1
}

func main() {
	arr := []int{1, 3, 5, 7, 9, 11, 13}
	dest := 6

	index := FirstMore(arr, dest)

	fmt.Println(arr[index])
}
