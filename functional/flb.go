package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)
// 定义 intGen 为 func() int 类型
type intGen func() int

/**
为函数实现接口
 */
func (g intGen) Read(
	p []byte) (n int, err error) {
	//panic("implement me")
	next := g()
	if (next > 10000) {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	return  strings.NewReader(s).Read(p)
}

/**
斐波拉切数列
 */
func fib() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a + b
		return a
	}
}

/**
打印 reader 内容
 */
func printFileContents(reader io.Reader)  {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
func main() {
	f := fib()
	//f.Read()
	/*for i := 0; i< 10 ; i++ {
		fmt.Println(f())
	}*/
	printFileContents(f)
}