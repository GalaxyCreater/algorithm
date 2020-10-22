package tree

import (
	"fmt"
)

/*
			head
	l1			|		r1
l2		r2		|	nil		rr
			r3


head l1 l2 r1 r3 r1 rr
*/
func CreateTree() (t *TreeNode) {
	t = nil
	t = &TreeNode{
		data:  "head",
		left:  nil,
		right: nil,
	}

	l1 := AddTreeNode(t, "l1", "l")
	r1 := AddTreeNode(t, "r1", "r")

	AddTreeNode(r1, "rr", "r")

	AddTreeNode(l1, "l2", "l")
	r2 := AddTreeNode(l1, "r2", "r")
	AddTreeNode(r2, "r3", "r")
	return
}

/*
*** 二叉树遍历
 */
// 前缀
func PrefixTravel(t *TreeNode) {
	if t == nil {
		//	fmt.Println("nil")
		return
	}

	fmt.Print(t.data, " ")
	PrefixTravel(t.left)
	PrefixTravel(t.right)
}

// 中缀
func MiddleTravel(t *TreeNode) {
	if t == nil {
		//	fmt.Println("nil")
		return
	}

	PrefixTravel(t.left)
	fmt.Print(t.data, " ")
	PrefixTravel(t.right)
}

// 后缀
func TailTravel(t *TreeNode) {
	if t == nil {
		//	fmt.Println("nil")
		return
	}

	PrefixTravel(t.left)
	PrefixTravel(t.right)
	fmt.Print(t.data, " ")
}

// 层级遍历
func LevelTravel(t *TreeNode) {
	lst := []*TreeNode{t}
	for len(lst) > 0 {
		ele := lst[0]
		if ele != nil {
			fmt.Print(ele.data, " ")
			if ele.left != nil {
				lst = append(lst, ele.left)
			}
			if ele.right != nil {
				lst = append(lst, ele.right)
			}
		}

		lst = lst[1:]
	}
}

/*
求树高度(递归方式)
*/
func TreeHeightRecursion(t *TreeNode) int {
	h := 0

	return _TreeHeightRecursion(t, h)
}
func _TreeHeightRecursion(t *TreeNode, h int) int {
	if t == nil {
		return 0
	}

	lh := _TreeHeightRecursion(t.left, h) + 1
	rh := _TreeHeightRecursion(t.right, h) + 1
	if lh > rh {
		return lh
	}
	return rh
}

/*
求树高度(非递归方式,利用遍历每层方式)
*/
func TreeHeight(t *TreeNode) int {
	if t == nil {
		return 0
	}

	lst := []*TreeNode{t}
	h := 0
	for len(lst) > 0 {
		h += 1
		length := len(lst)
		for i := 0; i < length; i++ {
			if lst[i].left != nil {
				lst = append(lst, lst[i].left)
			}
			if lst[i].right != nil {
				lst = append(lst, lst[i].right)
			}
		}

		lst = lst[length:]
	}

	return h
}
