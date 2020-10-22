package bserach

import (
	"fmt"
	"math"
	"testing"
)

func CheckFindResult(lst []int, f func(lst []int, val int) int) {
	result := true
	for i := 0; i < len(lst); i++ {
		ret := f(lst, lst[i])
		if ret == i {
			fmt.Printf("success find: %d [%d]%d\n", ret, i, lst[i])
		} else {
			result = false
			fmt.Printf("error find: %d [%d]%d\n", ret, i, lst[i])
		}
	}

	if result == false {
		fmt.Println("----------- FAILED -----------")
	} else {
		fmt.Println("----------- SUCCESS -----------")
	}
}

func TestBsearch(t *testing.T) {
	lst := []int{1, 2, 3, 4, 5, 6, 7}
	CheckFindResult(lst, Bsearch)
	lst = []int{1}
	CheckFindResult(lst, Bsearch)
	lst = []int{4}
	CheckFindResult(lst, Bsearch)
	lst = []int{1, 2, 4, 5, 6, 7}
	CheckFindResult(lst, Bsearch)
	lst = []int{1, 3, 4, 5, 6, 8, 11, 18}
	CheckFindResult(lst, Bsearch)
}

func TestBsearchRecursive(t *testing.T) {
	lst := []int{1, 2, 3, 4, 5, 6, 7}
	CheckFindResult(lst, BsearchRecursive)
	lst = []int{1}
	CheckFindResult(lst, BsearchRecursive)
	lst = []int{4}
	CheckFindResult(lst, BsearchRecursive)
	lst = []int{1, 2, 4, 5, 6, 7}
	CheckFindResult(lst, BsearchRecursive)
	lst = []int{1, 3, 4, 5, 6, 8, 11, 18}
	CheckFindResult(lst, BsearchRecursive)
}

func TestCalcSqrt(t *testing.T) {
	var v float64 = 3
	fmt.Printf("match ： %6f\n", math.Sqrt(v))
	fmt.Println("last:", CalcSqrt(v))

	v = 4
	fmt.Printf("match ： %6f\n", math.Sqrt(v))
	fmt.Println("last:", CalcSqrt(v))
}

func TestFindFirst(t *testing.T) {
	lst := []int{1, 3, 4, 5, 6, 8, 8, 8, 11, 18}
	fmt.Println(FindFirstEqual(lst, 8))
	fmt.Println(FindFirstEqualEx(lst, 8))
}

func TestFindLastEqual(t *testing.T) {
	lst := []int{1, 3, 4, 5, 6, 8, 8, 8, 11, 18}
	fmt.Println(FindLastEqual(lst, 8))
}

func TestFindFirstGreaterEqual(t *testing.T) {
	lst := []int{1, 3, 4, 5, 6, 8, 8, 8, 11, 18}
	fmt.Println(FindFirstGreaterEqual(lst, 8))
}

func TestFindLastLessEqual(t *testing.T) {
	lst := []int{1, 3, 4, 5, 6, 8, 8, 8, 11, 18}
	fmt.Println(FindLastLessEqual(lst, 8))
}

func TestBinsearchCircular(t *testing.T) {
	lst := []int{4, 5, 6, 1, 2, 3}
	CheckFindResult(lst, BinsearchCircular)
}
