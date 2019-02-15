package main

import "fmt"

// 简单排序算法

// 1.冒泡排序
func bubbleSort(arr []int) {
	var temp int
	size := len(arr)
	for i := 0; i < size-1; i++ {
		for j := 0; j < size-1-i; j++ {
			if arr[j] > arr[j+1] {
				temp = arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}
}

// 2.选择排序
func selectSort(arr []int) {
	var temp, min int
	size := len(arr)

	for i := 0; i < size-1; i++ {
		min = i
		for j := i; j < size; j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}

		if min != i {
			temp = arr[i]
			arr[i] = arr[min]
			arr[min] = temp
		}
	}

}

// 3.插入排序
func insertSort(arr []int) {
	var j int
	size := len(arr)

	for i := 1; i < size; i++ {
		for j = 0; j < i; j++ {
			if arr[j] < arr[i] {
				break
			}
		}
	}
}

func main() {
	arr := []int{9, 8, 7, 6, 5, 4, 2, 3, 1, 0}
	fmt.Println(arr)
	fmt.Println("======================")
	// bubbleSort(arr)
	// selectSort(arr)
	insertSort(arr)
	fmt.Println(arr)
}
