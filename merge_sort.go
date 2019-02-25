package main

import "fmt"

/**********************************************
 * 归并排序
 * 每次将数组从中间分成两部分
 * 直到两部分分别只有一个元素,或者只有一个元素
 * 这样每一个子数组都是有序的(只有一个元素)
 * 然后将两个有序的子数组合成一个有序的数组
 *********************************************/

func mergeSort(arr []int, start, end int) {
	if start >= end {
		return
	}

	mid := (start + end) / 2
	mergeSort(arr, start, mid)
	mergeSort(arr, mid+1, end)
	merge(arr, start, mid, end)
}

func merge(arr []int, start, mid, end int) {
	tmpArr := make([]int, end-start+1)

	i := start
	j := mid + 1
	k := 0
	for ; i <= mid && j <= end; k++ {
		if arr[i] < arr[j] {
			tmpArr[k] = arr[i]
			i++
		} else {
			tmpArr[k] = arr[j]
			j++
		}
	}

	for ; i <= mid; i++ {
		tmpArr[k] = arr[i]
		k++
	}
	for ; j <= end; j++ {
		tmpArr[k] = arr[j]
		k++
	}
	copy(arr[start:end+1], tmpArr)
}

func main() {
	arr := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	mergeSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
