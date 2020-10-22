package day_1_test

import (
	"errors"
	"fmt"
	"testing"
)

type ItemInterface interface {
	compare(rhs interface{}) bool
	GetValue() interface{}
	GetString() string
}

type IntItem struct {
	val int
}

func (it IntItem) compare(rhs interface{}) bool {
	if it.val > rhs.(int) {
		return true
	}

	return false
}

func (it IntItem) GetValue() interface{} {
	return it.val
}

func (it IntItem) GetString() string {
	return fmt.Sprintf("%d", it.val)
}

/*
有序数组实现
*/
type ArrayImpl struct {
	arr    []*ItemInterface
	length int
}

func CreateArry() *ArrayImpl {
	return &ArrayImpl{
		length: 0,
		arr:    make([]*ItemInterface, 2, 2),
	}
}

func (self *ArrayImpl) checkIndex(idx int) error {
	if idx >= len(self.arr) {
		return errors.New("out of range")
	}

	return nil
}

func (self *ArrayImpl) findIndexMore(val interface{}) int {
	i := 0
	for ; i < self.length; i++ {
		if (*self.arr[i]).compare(val) {
			return i
		}
	}

	return i
}

func (self *ArrayImpl) Insert(item ItemInterface, idx int) error {
	// 扩容
	if self.length >= len(self.arr) {
		new := make([]*ItemInterface, self.length*2, self.length*2)
		self.arr = append(self.arr, new...)
	}

	err := self.checkIndex(idx)
	if err != nil {
		return err
	}

	id := self.findIndexMore((item).GetValue())

	// 移动元素
	for i := self.length; i > id; i-- {
		self.arr[i] = self.arr[i-1]
	}
	self.arr[id] = &item
	self.length += 1

	return nil
}

func (self *ArrayImpl) PushBack(item ItemInterface) error {

	return self.Insert(item, self.length)
}

func (self *ArrayImpl) Find(item ItemInterface) (int, error) {
	for i := 0; i < self.length; i++ {
		it := *self.arr[i]
		if item.GetValue() == it.GetValue() {
			return i, nil
		}
	}

	return 0, errors.New("no value")
}

func (self *ArrayImpl) Delete(item ItemInterface) (data *ItemInterface, err error) {
	data = nil
	idx, err := self.Find(item)
	if err != nil {
		return nil, err
	}

	data = self.arr[idx]
	// 前移
	for i := idx; i < self.length; i++ {
		self.arr[idx] = self.arr[i+1]
	}
	self.length--

	return
}

func (self *ArrayImpl) Print() {
	for i := 0; i < self.length; i++ {
		tmp := self.arr[i]
		fmt.Printf("%v ", (*tmp).GetString())
	}
	fmt.Println("")
}

func TestArray(t *testing.T) {
	arr := CreateArry()
	arr.PushBack(IntItem{val: 2})
	arr.PushBack(IntItem{val: 1})
	arr.PushBack(IntItem{val: 4})
	arr.PushBack(IntItem{val: 3})

	arr.Print()

	arr.Delete(IntItem{val: 3})

	arr.Print()
}
