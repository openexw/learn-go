package main

import "fmt"

/**
包的变量
 */

 var (
 	a=1
 	b=3
 )
/**
变量申明
 */
func variable () {
	var num int
	var str string

	fmt.Printf("%d %q\n", num, str)
}

/**
变量赋值
 */
func variableSetValue() {
	var a,b int = 3, 4
	var s string = "qwr"
	fmt.Println(a, b, s)
}

/**
省略变量类型
 */
func variableTypeDeduction() {
	var a,b,c = 2, "d", true
	fmt.Println(a, b, c)
}

/**
省略var，记得 `=` 前面的 `:`
 */
func variableShorter() {
	a,b,c,d := 3, 4, true, "def"

	fmt.Println(a, b, c, d)
}
func main()  {
	fmt.Println("Hello")
	variable()

	variableSetValue()
	variableTypeDeduction()
	variableShorter()
}
