package tree

import "fmt"

//import "fmt"

type Book struct {
	title   string
	author  string
	subject string
	book_id int
}

/**
封装

函数名字的封装
首字母大写：public
首字母小写：private

正对与包

一个目录一个包

main包包包含可执行接口

*/
func main() {
	book := Book{title: "GO语言圣经", author: "nil", subject: "Go lang", book_id: 1}

	fmt.Println(book, "--")
}
