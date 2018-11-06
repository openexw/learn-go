package main

import (
	"bufio"
	"fmt"
	"os"
)

func tryDefer()  {
	defer fmt.Println(1)
	defer fmt.Println(2)
	// defer 在函数 return 返回之前执行
	fmt.Println(3)
	return
	//defer fmt.Println(4)
}

func writeFile(filename string)  {
	file, err := os.Create(filename)

	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	for i:=0; i<10; i++ {
		fmt.Fprintln(writer, i)
	}
}

func openFile (filename string) {
	_, err := os.Open(filename)
	if err != nil {
		if pathError, ok := err.(*os.PathError); !ok {
			fmt.Println(pathError.Op,
				pathError.Path,
				pathError.Err)
		}
	}
}
func main() {
	tryDefer() // 3 2 1


	writeFile("fib.txt")
}
