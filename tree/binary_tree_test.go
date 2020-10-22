package tree_test

import (
	"fmt"
	"test_code/algorithm/tree"
	"testing"
)

func TestTreeTravel(t *testing.T) {
	root := tree.CreateTree()
	tree.PrefixTravel(root)
	fmt.Println("")
	tree.MiddleTravel(root)
	fmt.Println("")
	tree.TailTravel(root)
	fmt.Println("")
	tree.LevelTravel(root)
	fmt.Println("")
}

func TestTreeHeight(t *testing.T) {
	root := tree.CreateTree()
	fmt.Println(tree.TreeHeightRecursion(root))
	fmt.Print("no recursion:", tree.TreeHeight(root), "\n")
}
