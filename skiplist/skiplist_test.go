package skiplist_test

import (
	"test_code/algorithm/skiplist"
	"testing"
)

func TestSkipList(t *testing.T) {
	sk := skiplist.ConsturctSkipList()
	sk.AddNode(&skiplist.StringData{
		Val: "a",
	}, 2)
	//	sk.AddNode("a", 2)
	// sk.AddNode("b", 1)
	// sk.AddNode("c", 3)
	sk.Print()

	// fmt.Println("--------find--------")
	// bp := sk.Find("b", 1)
	// ap := sk.Find("a", 2)
	// cp := sk.Find("c", 3)
	// fmt.Println(bp, unsafe.Pointer(bp))
	// fmt.Println(ap, unsafe.Pointer(ap))
	// fmt.Println(cp, unsafe.Pointer(cp))

	// sk.DelNode("a", 2)
	// fmt.Println("---------del--------")
	// sk.Print()
}
