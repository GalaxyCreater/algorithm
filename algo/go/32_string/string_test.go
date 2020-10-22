package _32_string

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	main := "abcacabcbcabcabc"
	pattern := "cabcab"
	//pattern := "babcab"

	fmt.Println(bmSearch(main, pattern))
}
