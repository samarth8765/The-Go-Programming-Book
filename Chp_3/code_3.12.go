package main

import (
	"fmt"
)

func main() {
	s := "abcda"
	t := "acbda"

	fmt.Println(checkAnagrams(s, t))
}

func checkAnagrams(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	a := make([]int, 256)
	b := make([]int, 256)

	for i := 0; i < len(s); i++ {
		a[s[i]]++
		b[t[i]]++
	}

	for i := 0; i < 255; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
