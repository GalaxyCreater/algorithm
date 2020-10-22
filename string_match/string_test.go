package string_match_test

import (
	"fmt"
	"test_code/algorithm/string_match"
	"testing"
)

func TestBmSearch(t *testing.T) {
	// main := "abcacabcbcabcabc"
	// //pattern := "cabcab"
	// pattern := "babcab"

	main := "防静电矢量发动机发电量附近"
	pattern := "发动机"
	fmt.Println(string_match.BmSearch(main, pattern))
}

/*
abcacabcbcabcabc
cabcab

0.
abcac‘a’ bcbcabcabc
cabca‘b’

1.
abc‘a’cab cbcabcabc
 ca‘b’cab

4.
abcacabcbc abcabc
	cabcab

6.
abcacab'c'bcab cabc
	  c'a'bcab

9.
abcacabcbcabcabc
	     cabcab
*/
