package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "中国人民共和国"		// UTF-8

	//for i,ch := range []byte(s) {	// 将字符对应的 utf8 对应编码输出
	/*for i,ch := range s {	// 将 utf8 转出成 unicode
		fmt.Printf("[%X %d]", ch, i)
	}*/

	//utf8.RuneCountInString(s) 获得字符数
	chs := []byte(s)
	for len(chs) > 0 {
		ch, size := utf8.DecodeRune(chs)
		chs = chs[size:]
		fmt.Printf("%c ", ch)
	}

	for i, ch := range []rune(s) {
		fmt.Printf("(%d, %c) ", i, ch)
	}

}
