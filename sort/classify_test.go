package sort

import (
	"fmt"
	"testing"
	"unicode"
)

func ClassifyEle(s string) string {

	numIdx := -1
	cIdx := -1

	lst := []rune(s)
	for i, v := range lst {
		if unicode.IsLetter(v) {
			if cIdx == -1 {
				cIdx = i
			}
			if numIdx != -1 {
				lst[i], lst[numIdx] = lst[numIdx], lst[i]
				numIdx++
			}
		} else {
			if numIdx == -1 {
				numIdx = i
			}
		}
	}

	return string(lst[:])
}

func TestClassifyEle(t *testing.T) {
	s := "a2b213fw23f"
	s = "1a2b213fw23f"
	fmt.Println(ClassifyEle(s))
}
