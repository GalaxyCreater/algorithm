package sort

import (
	"fmt"
	"testing"
)

func RunSort(f func(l []int)) {

	lst := []int{2, 1, 5, 6, 1, 4, 2, 4, 6, 3, 5, 3}
	fmt.Println("---------------")
	fmt.Println("source: ", lst)
	f(lst)
	fmt.Println("result: ", lst)

	lst = []int{5, 3, 1, 0}
	fmt.Println("---------------")
	fmt.Println("source: ", lst)
	f(lst)
	fmt.Println("result: ", lst)

	lst = []int{5, 1}
	fmt.Println("---------------")
	fmt.Println("source: ", lst)
	f(lst)
	fmt.Println("result: ", lst)
}

func TestBubbleSort(t *testing.T) {
	RunSort(BubbleSort)
}

func TestInsertionSort(t *testing.T) {
	RunSort(InsertionSort)
}

func TestMergeSort(t *testing.T) {
	RunSort(MergeSort)
}

func TestQuickSort(t *testing.T) {
	RunSort(QuickSort)
}

func TestFindTop(t *testing.T) {
	lst := []int{2, 1, 5, 6, 1, 4, 2, 4, 6, 3, 5, 3}
	fmt.Println(lst)
	fmt.Println(FindTopK(lst, 3))
	fmt.Println(findTopKLoop(lst, 3))
}

func TestBucketSort(t *testing.T) {
	RunSort(BucketSort)
}

func TestCountingSort(t *testing.T) {
	RunSort(CountingSort)
}

func TestRadixSort(t *testing.T) {
	lst := []int{32, 43, 12, 44, 44, 65, 21, 1}
	RadixSort(lst)
	fmt.Println(lst)
}

func TestMyBubbleSort(t *testing.T) {
	lst := []int{32, 43, 12, 44, 44, 65, 21, 1}
	MyBubbleSort(lst)
	fmt.Println(lst)
}

func TestMyQuickSort(t *testing.T) {
	lst := []int{6, 11, 3, 9, 8}
	MyQuickSort(lst)
	fmt.Println(lst)
}
