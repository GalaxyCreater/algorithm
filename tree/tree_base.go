package tree

type TreeNode struct {
	data  string
	left  *TreeNode
	right *TreeNode
}

func AddTreeNode(t *TreeNode, data string, chType string) *TreeNode {
	node := &TreeNode{
		data:  data,
		left:  nil,
		right: nil,
	}
	if chType == "r" {
		t.right = node
	} else {
		t.left = node
	}

	return node
}
