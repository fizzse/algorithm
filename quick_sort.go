package main

import "fmt"

/**********************************************
 * 快速排序
 * 每次随机找一个数字
 * 每次排序保证左边都比它小
 * 每次排序保证右边都比它大
 * 直到拆分到只有一个元素
 *********************************************/

func QuickSort(arr []int) {
	SeparateSort(arr, 0, len(arr)-1)
}

func SeparateSort(arr []int, start, end int) {
	if start >= end {
		return
	}

	i := Partition(arr, start, end)

	SeparateSort(arr, start, i-1)
	SeparateSort(arr, i+1, end)
}

func Partition(arr []int, start, end int) int {
	pivot := arr[end]

	i := start
	for j := start; j < end; j++ {
		if arr[j] < pivot {
			if i == j {
				arr[i], arr[j] = arr[j], arr[i]
			}
			i++
		}
	}

	arr[i], arr[end] = arr[end], arr[i]
	return i
}

func main() {
	arr := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	QuickSort(arr)
	fmt.Println(arr)
}
