package main

import "fmt"

/**********************************************
 * 合并两个有序数组
 *********************************************/

func ArrayMerge(a []int, b []int) []int {
	s1 := len(a)
	s2 := len(b)
	size := s1 + s2
	temp := make([]int, size)
	j, k := 0, 0
	for i := 0; i < size; i++ {
		if j < s1 && k < s2 {
			if a[j] < b[k] {
				temp[i] = a[j]
				j++
			} else {
				temp[i] = b[k]
				k++
			}
		} else if j == s1 {
			temp[i] = b[k]
			k++
		} else if k == s2 {
			temp[i] = a[j]
			j++
		}
	}

	return temp
}

func main() {
	a := []int{1, 3, 5, 7, 9}
	b := []int{0, 2, 4, 6, 8}
	merge := ArrayMerge(a, b)
	fmt.Println(merge)
}
