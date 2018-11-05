package main

import "fmt"

/**
ababb => ab
bbbbb => b
*/
func getNoRepearStr(s string) int {
	last0current := make(map[byte]int)
	start := 0
	maxLen := 0

	for i, ch := range []byte(s) {
		if lastI, ok := last0current[ch]; ok && lastI >= start {
			start = lastI + 1
		}

		if i-start+1 > maxLen {
			maxLen = i - start + 1
		}
		last0current[ch] = i
	}
	return maxLen
}
func main() {
	fmt.Println(getNoRepearStr("abcasd"))
}
