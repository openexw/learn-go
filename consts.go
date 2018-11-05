package main

func consts() {
	const (
		FILENAME = "a.txt"
		a, b     = 3, 5
	)
	const m = 23
	println(FILENAME) // a.txt
}

func enums() {
	const (
		a = 1
		d = 2
		c = 4
	)
	println(a, d, c)
	const (
		// iota 作为一个自增值
		cpp = iota
		java
		python
		golang
	)

	println(cpp, java, python, golang) // 0 1 2 3

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	// 1 1024 1048576 1073741824 1099511627776 1125899906842624
	println(b, kb, mb, gb, tb, pb)
}
func main() {
	consts()
	enums()
}
