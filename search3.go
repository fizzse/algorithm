package main

import "fmt"

/* 一堆人围在依次数1,2,3,1,2,3
 * 数到3的退出游戏
 * 剩下的人继续数,直到剩下最后一个人
 */

type Person struct {
	current int
	origin  int
}

// 一次游戏
func game(count int, arr []*Person) int {
	// 累计编号
	seq := 0

	// 一轮游戏
	for i := 0; i < len(arr); i++ {
		// 排除无效人员
		if arr[i] == nil {
			continue
		}

		seq = arr[i].current
		num := arr[i].current % 3
		if num == 0 {
			num = 3
		}
		fmt.Printf("person current: %d origin: %d say: %d\n", arr[i].current, arr[i].origin, num)
		if num == 3 {
			fmt.Printf("person current: %d origin: %d out of game !!\n", arr[i].current, arr[i].origin)
			arr[i] = nil
			count--
		}
	}

	fmt.Println("last person count: ", count)

	// 剩下的人重新编号
	for k := 0; k < len(arr); k++ {
		// 去除无效人员
		if arr[k] == nil {
			continue
		}

		// 重新编号，在上一轮的基础上累加
		arr[k].current = seq + 1
		seq++
		fmt.Printf("person origin: %d in the new game current: %d \n", arr[k].origin, arr[k].current)
	}

	return count
}

// 循环比赛
func loop(count int, arr []*Person) {
	// 退出条件为还剩一个人
	j := 1
	for count > 1 {
		fmt.Printf("第%d轮游戏开始\n", j)
		count = game(count, arr)
		j++

		fmt.Println("==========================================================")
		fmt.Println()
		fmt.Println()
	}

	// 得到最后剩下的人
	for i := 0; i < len(arr); i++ {
		if arr[i] == nil {
			continue
		}

		fmt.Printf("person origin:%d is alive\n", arr[i].origin)
	}
}

// 方法2(简单)
func foo(num int) {
	arr := make([]int, num)
	for i := 0; i < num; i++ {
		arr[i] = i + 1
	}

	// 每次淘汰1个人,循环num-1次即最后剩下1个人
	for i := 0; i < num-1; i++ {
		// 数到1的人放到队尾
		arr = append(arr[1:], arr[0])
		// 数到2的人放到队尾
		arr = append(arr[1:], arr[0])
		// 数到3的人淘汰
		arr = arr[1:]
	}

	// 最后剩下的人
	fmt.Println(arr)
}

func main() {
	num := 30
	// 方法1
	arr := make([]*Person, num)

	for i := 0; i < num; i++ {
		arr[i] = &Person{current: i + 1, origin: i + 1}
	}

	loop(num, arr)

	// 方法2
	foo(num)
}
