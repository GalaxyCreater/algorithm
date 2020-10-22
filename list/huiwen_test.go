package list

import (
	"fmt"
	"strconv"
	"testing"
)

/*
判断一个列表是不是回文，回文： "abcba"
*/

func CheckHuiWen(self *Mylist) bool {
	lhs := []int{}

	p1 := self.head.next
	p2 := self.head.next
	for p2 != nil {
		// 走两步
		if p2.next == nil {
			p1 = p1.next
			break
		} else {
			p2 = p2.next
		}
		p2 = p2.next

		// 走一步
		lhs = append(lhs, p1.data)
		p1 = p1.next
	}

	for i := len(lhs) - 1; i >= 0; i-- {
		v := lhs[i]
		if p1 == nil {
			return false
		}
		if v != p1.data {
			return false
		}
		p1 = p1.next
	}

	return true
}

func TestHuiWen(t *testing.T) {
	//arr := []int{1, 2, 2, 2, 1}
	arr := []int{1, 2, 3, 3, 2, 1}
	obj := CreateList(arr)
	obj.Print()
	flag := CheckHuiWen(obj)
	if flag == false {
		fmt.Println("not ....")
	} else {
		fmt.Println("is ...")
	}

	strconv.Atoi
}

func Ai(s string) int {
	n := 0
	for _, ch := range []byte(s) {
		ch -= '0'
		if ch > 9 {
			return 0, &NumError{fnAtoi, s0, ErrSyntax}
		}
		n = n*10 + int(ch)
	}

	return n
}

func TestHuiWen(t *testing.T) {
	n := Ai("1003")
	fmt.Println(n)
}
