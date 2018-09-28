/*
 * 求一个正整数的补码中包含多少个1
 *
 */

package main

import (
	"fmt"
	"unsafe"
)

// 位移法
func foo1(n uint) int {
	ret := 0
	len := unsafe.Sizeof(n)
	for i := 0; i < int(len); i++ {
		if (n & 1) == 1 {
			ret++
		}
		n >>= 1
	}

	return ret
}

// 求与法
// FIXME n & n-1 能去掉最后面的那个1,2的整数次幂只有1个1,如果n&(n-1)=0,n是2的整数次幂
func foo2(n uint) int {
	ret := 0
	for n != 0 {
		ret++
		n &= (n - 1)
	}

	return ret
}

// 查表法
func foo3(n uint) int {
	// 一个uint32的正整数n，一旦n的值确定，n的二进制表示中包含多少个1也就确定了，理论上无需重新计算
	// 1的二进制表示中包含1个1
	// 2的二进制表示中包含1个1
	// 3的二进制表示中包含2个1
	// ...
	// 58585858的二进制表示中包含15个1
	// 提前计算好结果数组：
	// result[1]=1;
	// result[2]=1;
	// result[3]=2;
	// ...
	// result[58585858]=15;
	// return result[n];
	// 查表法的好处是，时间复杂度为O(1)，潜在的问题是，需要很大的内存
	// 假如被分析的整数是uint32，打表数组需要记录2^32个正整数的结果,n的二进制表示最多包含32个1，存储结果的计数，使用5个bit即可
	// 故，共需要内存2^32 * 5bit = 2.5GB
	return 0
}

// 二次查表法
func foo3(n uint) int {
	// 查表法，非常快，只查询一次，但消耗内存太大，在工程中几乎不被使用
	// 算法设计，本身是一个时间复杂度与空间复杂度的折衷，增加计算次数，往往能够减少存储空间
	// 思路：
	// （1）把uint32的正整数n，分解为低16位正整数n1，和高16正整数n2；
	// （2）n1查一次表，其二进制表示包含a个1；
	// （3）n2查一次表，其二进制表示包含b个1；
	// （4）则，n的二进制表示包含a+b个1；
	// uint16 n1 = n & 0xFFFF;
	// uint16 n2 = (n>>16) & 0xFFFF;
	// return  result[n1]+result[n2];
	// 被分析的整数变成uint16，打表数组需要记录2^16个正整数的结果。
	// n1和n2的二进制表示最多包含16个1，存储结果的计数，使用4个bit即可。
	// 故，共需要内存2^16 * 4bit = 32KB。
	// 计算量多了1次（1倍），内存占用量却由2.5G降到了32K（1万多倍），是不是很有意思？
	return 0
}

func main() {
	var num uint
	fmt.Println("请输入一个数")
	fmt.Scan(&num)
	fmt.Printf("%d 的补码包含 [%d] 个1\n", num, foo1(num))
	fmt.Printf("%d 的补码包含 [%d] 个1\n", num, foo2(num))
}
