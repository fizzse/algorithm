/*
 * 手写单向链表包含以下api
 * Init
 * PushFront
 * PushBack
 * Foreach
 * Foreach
 * Foreach
 */
package main

import (
	"fmt"
)

type Node struct {
	Data interface{}
	Next *Node
}

type List struct {
	Header *Node
	Size   int
}

func (l *List) Init() {
	l.Header = &Node{Data: 1, Next: nil}
	l.Size = 0
}

func (l *List) PushFront(data interface{}) {
	node := &Node{Data: data}
	node.Next = l.Header.Next
	l.Header.Next = node
	l.Size++
}

func (l *List) PushBack(data interface{}) {
	node := &Node{Data: data}
	tmp := l.Header
	for i := 0; i < l.Size; i++ {
		tmp = tmp.Next
	}

	tmp.Next = node
	l.Size++
}

func (l *List) Foreach() {
	tmp := l.Header.Next
	for i := 0; i < l.Size; i++ {
		fmt.Printf("%d\t", tmp.Data)
		tmp = tmp.Next
	}

	fmt.Println()
}

// 头插法,采用辅助指针，将链表中的每一个节点进行 PushFront
func (l *List) ReverseByPushFront() {
	if l.Size < 2 {
		return
	}

	p := l.Header.Next
	l.Header.Next = nil
	for i := 0; i < l.Size; i++ {
		tmp := p
		p = p.Next
		tmp.Next = l.Header.Next
		l.Header.Next = tmp
	}
}

// 就地逆置非递归实现,将下一个节点设置为每一次的新的头结点
func (l *List) Reverse() {
	if l.Size < 2 {
		return
	}

	p := l.Header.Next
	var newH *Node

	for i := 0; i < l.Size; i++ {
		tmp := p.Next
		p.Next = newH
		newH = p
		p = tmp
	}

	l.Header.Next = newH
}

// 就地逆置递归实现
func (l *List) inReverse(head *Node) *Node {
	// 如果没有节点，或者只有一个节点
	if head == nil || head.Next == nil {
		return head
	}

	newH := l.inReverse(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newH
}

func (l *List) ReverseByRecursive() {
	head := l.Header.Next
	newH := l.inReverse(head)
	l.Header.Next = newH
}

func main() {
	l := List{}
	l.Init()

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.PushBack(4)
	l.Foreach()
	l.ReverseByPushFront()
	l.Foreach()
	l.Reverse()
	l.Foreach()
	l.ReverseByRecursive()
	l.Foreach()
}
