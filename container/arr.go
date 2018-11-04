package main

func arr() {
	// 数组的申明
	var n [10]int;
	var i,j int

	// 数组赋初值
	for i = 0; i < 10; i++ {
		n[i] = i + 10
	}

	// 数组遍历
	for j=0; j < 10 ; j++ {
		println(n[j])
	}

	// len() 方法
	for j=0; j < len(n) ; j++ {
		println(n[j])
	}
	
	// range 关键字
	for j := range n {
		println(n[j])
	}
	for j,v := range n {
		println(j, v)	// j 下标，v 值
	}
	for _,v := range n {
		println(v)	// 省略j
	}
	//var m = [...]int{1, 3, 2}
}

func printArr(arr []int)  {
	
}
func main() {
	//arr()
}
