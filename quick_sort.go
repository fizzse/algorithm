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
	// 取最后一个元素为判断数据
	// 保证右边的都比自己小
	// 左边(没有元素),都比自己大
	pivot := arr[end]

	i := start
	for j := start; j < end; j++ {
		// 如果右边的数据小,则哨兵后移
		if arr[j] < pivot {
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
	arr := []int{1, 3, 5, 7, 9, 2, 4, 6, 8, 0}
	QuickSort(arr)
	fmt.Println(arr)
}
