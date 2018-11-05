package main

import "fmt"

func slice() {
	arr := [...]int{1, 2, 3, 4, 5, 6}
	//s := arr[2:4]
	//println("sds", arr[2:])
	fmt.Println("as", arr[3:])
}
func main() {
	slice()

	//var s []int	// slice
	var num []int

	// 追加空元素
	num = append(num, 0)
	fmt.Println(num) // [0]

	// 追加一个元素
	num = append(num, 1)
	fmt.Println(num) // [0 1]

	// 追加多个元素
	num = append(num, 2, 3, 4)
	fmt.Println(num, cap(num)) // [0 1 2 3 4] 6

	// 创建切片 num1 是之前切片的两倍容量
	num1 := make([]int, len(num), cap(num)*2)
	fmt.Println(num1, cap(num1)) // [0 0 0 0 0] 12

	copy(num1, num)
	fmt.Println(num1) // [0 1 2 3 4]
}
