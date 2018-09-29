/*
 * 斐波那契数列
 * f(0) = 0
 * f(1) = 1
 * f(n) = f(n-1) + f(n-2)
 *
 */

package main

import "fmt"

// 1.递归法
// 斐波那契数列求解，如果用直接法，时间复杂度是指数级的，不可行
// 如果没有太大的把握，工程中尽量少使用递归，容易把自己玩死
func foo1(n uint) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	return foo1(n-1) + foo1(n-2)
}

// 2.正推法
// 时间复杂度O(n)
func foo2(n uint) uint {
	arr := make([]uint, n+1)
	arr[0] = 0
	arr[1] = 1
	for i := 2; i <= int(n); i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}

	return arr[n]
}

// 3.通项公式法，数学不好的去百度
// f(n)=(1/√5)*{[(1+√5)/2]^n -[(1-√5)/2]^n}
// 套公式即可

// 4.查表法，空间换时间的老套路
// 将提前算好的结果放在数组中
// f(0) = 0
// f(1) = 1
// f(2) = 1
// f(3) = 2
// ....

func main() {
	n := uint(10)
	fmt.Printf("f(%d) = %d\n", n, foo1(n))
	fmt.Printf("f(%d) = %d\n", n, foo2(n))
}
