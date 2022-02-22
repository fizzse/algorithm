package main

import "fmt"

/**********************************************
 * 快速排序
 * 每次随机找一个数字
 * 每次排序保证左边都比它小
 * 每次排序保证右边都比它大
 * 直到拆分到只有一个元素
 *********************************************/

func quickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	i := partition(arr)
	quickSort(arr[:i])
	quickSort(arr[i:])
}

func partition(arr []int) int {
	end := len(arr) - 1
	pivot := arr[end]
	i := 0

	for j := 0; j < end; j++ {
		if arr[j] <= pivot {
			if i != j {
				arr[i], arr[j] = arr[j], arr[i]
			}
			i++
		}
	}

	arr[i], arr[end] = arr[end], arr[i]
	return i
}

func main() {
	arr := []int{1, 3, 5, 7, 0, 2, 4, 6, 8, 9, 1}
	quickSort(arr)
	fmt.Println(arr)
}

