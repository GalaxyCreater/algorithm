package _7_skiplist

import (
	"fmt"
	"testing"
)

func TestSkipList(t *testing.T) {
	sl := NewSkipList()

	sl.Insert("leo", 95)
	fmt.Println(sl.head.forwards[0])
	fmt.Println(sl.head.forwards[0].forwards[0])
	fmt.Println(sl)
	fmt.Println("-----------------------------")

	sl.Insert("jack", 88)
	fmt.Println(sl.head.forwards[0])
	fmt.Println(sl.head.forwards[0].forwards[0])
	fmt.Println(sl.head.forwards[0].forwards[0].forwards[0])
	fmt.Println(sl)
	fmt.Println("-----------------------------")

	sl.Insert("lily", 100)
	fmt.Println(sl.head.forwards[0])
	fmt.Println(sl.head.forwards[0].forwards[0])
	fmt.Println(sl.head.forwards[0].forwards[0].forwards[0])
	fmt.Println(sl.head.forwards[0].forwards[0].forwards[0].forwards[0])
	fmt.Println(sl)
	fmt.Println("-----------------------------")

	fmt.Println(sl.Find("jack", 88))
	fmt.Println("-----------------------------")

	sl.Delete("leo", 95)
	fmt.Println(sl.head.forwards[0])
	fmt.Println(sl.head.forwards[0].forwards[0])
	fmt.Println(sl.head.forwards[0].forwards[0].forwards[0])
	fmt.Println(sl)
	fmt.Println("-----------------------------")
}
