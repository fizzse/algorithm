package main

import (
	"fmt"
)

type Node struct {
	Val  int
	Next *Node
}

func foreachList(head *Node) {
	p := head
	for p != nil {
		fmt.Printf("%d\t", p.Val)
		p = p.Next
	}

	fmt.Println()
}

// 获取链表中倒数第n个元素
func findKthToTail(head *Node, k int) *Node {
	p := head
	q := head
	i := 0

	// 总长度为n,要查找倒数第k个元素,正序需要执行n-k次
	// 也就是当循环因子等于k的时候，从头指针开始走,全部执行需要n次
	// 当i=k次,跟着遍历整体链表一起开始跑,一共执行n-k次,得到倒数第k个元素
	for i = 0; p != nil; i++ {
		p = p.Next
		if i >= k {
			q = q.Next
		}
	}

	if i < k {
		return nil
	}

	return q
}

func main() {
	pos := 0

	head := &Node{Val: 0}
	n1 := &Node{Val: 1}
	n2 := &Node{Val: 2}
	n3 := &Node{Val: 3}
	n4 := &Node{Val: 4}
	n5 := &Node{Val: 5}
	head.Next = n1
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5

	fmt.Println("链表所有元素")
	foreachList(head)
	fmt.Println("---------------------------------------------------------------------")

	fmt.Println("请输入要获取倒数第几个元素")
	fmt.Scan(&pos)
	ret := findKthToTail(head, pos)
	fmt.Printf("倒数第%d个元素为:[%d]\n", pos, ret.Val)
}
