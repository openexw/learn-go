package main

// 加法
func add(a, b int) int {
	return (a + b)* 2 - a - b
}

// 减法
func reduce(a,b int) int  {
	return  a - b
}

// 获取不重复字符串的数字
func getNoRepeatStr(s string) int {
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
	println(add(1, 2))
	println(reduce(10, 2))
}
