package main

/**
两个值进行交换
 */
func swap(a, b *int)  {
	*a, *b = *b, *a
}
func main() {
	var a int = 2
	// pa是a的指针 & 取地址符号
	var pa *int = &a	// 申明指针变量并赋初值 。*pa表示的是变量 pa； 表示的是变量的存储地址
	//*pa = 3
	println(pa)	// 2
	a, b := 3, 5
	swap(&a, &b);
	println(a, b)
}