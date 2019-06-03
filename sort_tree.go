package main

import "fmt"

type Node struct {
	Data   int
	LChild *Node
	RChild *Node
}

type Tree struct {
	Root *Node
}

func NewTree(root int) *Tree {
	return &Tree{
		Root: &Node{Data: root},
	}
}

// 实现二叉查找树(排序树)的插入与遍历

// 遍历树
func (t *Tree) Foreach() {
	p := t.Root
	if p == nil {
		fmt.Println("empty tree not can foreach")
		return
	}

	foreach(p)
	fmt.Println()
}

// 中序遍历
func foreach(root *Node) {
	if root == nil {
		return
	}

	foreach(root.LChild)
	fmt.Printf("%d\t", root.Data)
	foreach(root.RChild)
}

// 插入
func (t *Tree) Insert(data int) {
	node := &Node{Data: data}
	p := t.Root
	if p == nil {
		fmt.Println("empty tree not can insert")
		t.Root = node
		return
	}

	for p != nil {
		if p.Data > data {
			if p.LChild == nil {
				p.LChild = node
				return
			}

			p = p.LChild
		} else {
			if p.RChild == nil {
				p.RChild = node
				return
			}

			p = p.RChild
		}
	}
}

// 查找
func (t *Tree) GetNode(data int) *Node {
	p := t.Root
	if p == nil {
		fmt.Println("empty tree not can search")
		return nil
	}

	for p != nil {
		if p.Data > data {
			p = p.LChild
		} else if p.Data < data {
			p = p.RChild
		} else {
			return p
		}
	}

	return nil
}

func main() {
	t := NewTree(3)
	t.Insert(1)
	t.Insert(10)
	t.Insert(5)
	t.Insert(7)
	t.Insert(8)
	t.Insert(9)
	t.Insert(2)
	t.Insert(4)
	t.Insert(5)
	t.Insert(6)

	t.Foreach()
}
