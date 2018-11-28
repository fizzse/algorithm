package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 根据一棵树的先序遍历结果与中序遍历结果还原树结构
// 根据先序与后序无法确定树结构

// 在数组中获取目前元素的下标
func getIndex(src []int, target int) int {
	for i := 0; i < len(src); i++ {
		if src[i] == target {
			return i
		}
	}

	return 0
}

// 先序遍历
func preForeach(root *TreeNode) {
	if root == nil {
		return
	}

	// 先序遍历 根--左--右
	fmt.Printf("%d\t", root.Val)
	preForeach(root.Left)
	preForeach(root.Right)
}

// 中序遍历
func midForeach(root *TreeNode) {
	if root == nil {
		return
	}

	// 中序遍历 左--根--右
	midForeach(root.Left)
	fmt.Printf("%d\t", root.Val)
	midForeach(root.Right)
}

// 后序遍历
func postForeach(root *TreeNode) {
	if root == nil {
		return
	}

	// 先序遍历 左--右--根
	postForeach(root.Left)
	postForeach(root.Right)
	fmt.Printf("%d\t", root.Val)
}

func arrayForeach(array []int) {
	for i := 0; i < len(array); i++ {
		fmt.Printf("%d\t", array[i])
	}

	fmt.Println()
}

// 重建树结构
func reconstruction(pre, mid []int) *TreeNode {
	length := len(pre)
	// 递归退出条件
	if length == 0 {
		return nil
	}

	// 参数错误
	if length != len(mid) {
		return nil
	}

	// 先序的第一个元素就是根
	root := &TreeNode{}
	root.Val = pre[0]
	// 获取中序中根的位置
	index := getIndex(mid, pre[0])

	// 中序遍历中根的左边全是左子树的元素
	midLeft := append(mid[0:index])
	// 中序遍历中根的右边全是右子树的元素
	midRight := append(mid[index+1:])

	// 先序遍历中,根的右边到第index+1元素都是左子树元素
	preLeft := append(pre[1 : index+1])
	// 先序遍历中,根的第index+1到最后元素都是左子树元素
	preRight := append(pre[index+1:])

	// 递归获取左根和右根
	root.Left = reconstruction(preLeft, midLeft)
	root.Right = reconstruction(preRight, midRight)

	return root
}

func main() {
	pre := []int{1, 2, 4, 7, 3, 5, 6, 8}
	mid := []int{4, 7, 2, 1, 5, 3, 8, 6}

	root := reconstruction(pre, mid)

	// 先序检查结果
	fmt.Println("---------------------------------------------------------------------")
	fmt.Println("先序遍历")
	arrayForeach(pre)
	preForeach(root)
	fmt.Println()

	// 先序检查结果
	fmt.Println("---------------------------------------------------------------------")
	fmt.Println("中序遍历")
	arrayForeach(mid)
	midForeach(root)
	fmt.Println()

	// 打印后序遍历
	fmt.Println("---------------------------------------------------------------------")
	fmt.Println("后序遍历")
	postForeach(root)
	fmt.Println()
}
