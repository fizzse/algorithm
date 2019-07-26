package main

/**********************************
 * 堆和堆排序
 * 本例中所展示的为大顶堆
 * 堆是完全二叉树采用数组法表示
 *********************************/

import "fmt"

type Heap struct {
	data []int
	cap  int
	size int
}

// 初始化堆容量
func InitHeap(cap int) *Heap {
	return &Heap{
		data: make([]int, cap+1),
		cap:  cap,
		size: 0,
	}
}

// 插入数据
func (h *Heap) Insert(data int) {
	if h.size == h.cap {
		return
	}

	h.size++

	index := h.size
	h.data[index] = data

	for index/2 > 0 && h.data[index] > h.data[index/2] {
		h.data[index], h.data[index/2] = h.data[index/2], h.data[index]
		index = index / 2
	}
}

// 删除堆顶元素
func (h *Heap) DeleteTop() {
	if h.size == 0 {
		return
	}

	h.data[1] = h.data[h.size]
	h.size--

	h.Heapify(1, h.size)
}

// 堆化
func (h *Heap) Heapify(start, end int) {
	for {
		maxpos := start
		if start*2 <= end && h.data[start] < h.data[start*2] {
			maxpos = start * 2
		}

		if start*2+1 <= end && h.data[maxpos] < h.data[start*2+1] {
			maxpos = start*2 + 1
		}

		if maxpos == start {
			break
		}

		h.data[start], h.data[maxpos] = h.data[maxpos], h.data[start]
		start = maxpos
	}
}

func main() {
	h := InitHeap(10)
	h.Insert(1)
	h.Insert(3)
	h.Insert(5)
	h.Insert(7)
	h.Insert(9)
	h.Insert(2)
	h.Insert(4)
	h.Insert(6)
	h.Insert(8)
	h.Insert(10)

	fmt.Println(h.data[1:])

	h.DeleteTop()
	fmt.Println(h.data[1:])
}
