package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

/**
除法运算
@test 返回多个值
@q 商
@r 余数
*/
func div(a, b int) (q, r int) {
	return a / b, a % b
}

func apply(op func(int, int) int, a, b int) int {
	pointer := reflect.ValueOf(op).Pointer()
	name := runtime.FuncForPC(pointer).Name()
	fmt.Printf("Calling function %d with args "+"(%d, %d)", name, a, b)
	return op(a, b)
}
func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}

	return s
}
func main() {
	//a,b := div(13, 3)
	//println(a, b)	// 4 1

	println(apply(pow, 3, 4))

	println(apply(
		// 匿名函数
		func(a int, b int) int {
			return int(math.Pow(float64(a), float64(b)))
		}, 3, 4))

	println(sum(1, 2, 3, 4, 5))
}
