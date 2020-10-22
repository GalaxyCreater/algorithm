package day_1_test

import "container/list"

type SigNode struct {
	next *SigNode
	val  int
}

type SingleList struct {
	head *SigNode
}

func CreateSingleList() *SingleList {
	return &SingleList{
		head: &SigNode{next: nil, val: 0},
	}
}

func (self *SingleList) InsertBefore(v int) {
	// n := &SigNode{
	// 	val:  v,
	// 	next: nil,
	// }

	l := list.New()
	l.Back()
}
