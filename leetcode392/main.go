package main

import (
	"fmt"
)

func main() {
	fmt.Println("Output: ", isSubsequence("abc", "ahbgdc"))
}

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	if len(t) == 0 {
		return false
	}
	counter := 0
	for i := 0; i < len(t); i++ {

		if counter == len(s) {
			return true
		}
		if t[i] == s[counter] {
			counter++
		}

	}
	if counter == len(s) {
		return true
	}

	return false

}
