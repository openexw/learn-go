package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/**
十进制转二进制
*/
func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

/**
每读一行文件
*/
func readFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		println(scanner.Text())
	}
}

/**
九九乘法表
*/
func multiplication() {
	for i := 1; i <= 9; i++ { // i 控制行，以及计算的最大值
		for j := 1; j <= i; j++ { // j 控制每行的计算个数
			fmt.Printf("%d*%d=%d ", j, i, j*i)
		}
		fmt.Println("")
	}
}
func main() {
	println(
		convertToBin(5),  // 101
		convertToBin(10), // 1010
		convertToBin(0),  // ''
	)

	//readFile("a.txt")

	multiplication()
}
