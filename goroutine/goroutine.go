package main

import (
	"runtime"
	"time"
)

func main() {
	for i := 0; i < 100; i++ {
		// 并发执行这个函数
		go func(i int) {
			for  {
				//fmt.Printf("hello goroutine %d\n", i)
				// 手动交出控制权，一般不会用到
				runtime.Gosched()
			}
		}(i)

		time.Sleep(time.Millisecond)

	}
}