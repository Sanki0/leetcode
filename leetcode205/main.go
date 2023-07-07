package main

import (
	"fmt"
)

func main() {
	fmt.Println("Output: ", isIsomorphic("egg", "add"))
}

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	s2t := make(map[byte]byte)
	t2s := make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		if s2t[s[i]] == 0 {
			s2t[s[i]] = t[i]
		} else if s2t[s[i]] != t[i] {
			return false
		}
		if t2s[t[i]] == 0 {
			t2s[t[i]] = s[i]
		} else if t2s[t[i]] != s[i] {
			return false
		}
	}
	return true
}
