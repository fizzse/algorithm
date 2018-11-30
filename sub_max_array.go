package main

import (
	"fmt"
)

// 在一个整数数组中有正数有负数,获取和最大的连续的子序列
// 1.固定长度,连续子序列的长度都一样,求最大的那个
func findMaxSubArrayFixed(array []int, length int) (int, []int) {
	size := len(array)
	if size == 0 {
		return 0, nil
	}

	//FIXME 子序列长度小于原序列
	if length > size {
		return 0, nil
	}

	max := 0
	index := 0
	// 一共有size-length的连续子序列
	for i := 0; i < size-length+1; i++ {
		// 对子序列求和
		sum := 0
		for j := i; j < i+length; j++ {
			sum += array[j]
		}

		// 获取最大的和
		if max < sum {
			max = sum
			index = i
		}
	}

	return max, array[index : index+length]
}

// 2.长度不固定,最少长度为1,求最大的子序列
func findMaxSubArray(array []int) (int, []int) {
	size := len(array)
	if size == 0 {
		return 0, nil
	}

	max, sum := 0, 0
	start, end := 0, 0
	for i := 0; i < size; i++ {
		// 如果前边的结果小于0,则加上后边的数,肯定会变小，所以前边的全部舍弃
		// 最后一次需要舍弃前边的序列的时候,此时的下标为最大序列的开始索引
		if sum < 0 {
			sum = array[i]
			start = i
		} else {
			sum += array[i]
		}

		// 如果后边的数是正数,和会变大,最大值则为当前的和
		// 当后边的数的和小于0时,此时的下标为最大序列的结束索引
		if max < sum {
			end = i
			max = sum
		}
	}

	return max, array[start : end+1]
}

func main() {
	src := []int{1, -10, 5, 6, -12, 20, -7, 8, 25, -10, 22, 9, -40}
	length := 10
	max, sub := findMaxSubArrayFixed(src, length)
	fmt.Printf("长度为%d的最大子序列的和为:[%d],子序列为:\n", length, max)
	fmt.Println(sub)

	fmt.Println("----------------------------------------------------------------")

	max, sub = findMaxSubArray(src)
	fmt.Printf("没有长度限制的最大子序列的和为:[%d],子序列为:\n", max)
	fmt.Println(sub)
}
