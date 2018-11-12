package main

import (
	"fmt"
	"time"
)

// create worker
// send data
func createWorker(id int)  chan<- int{
	c := make(chan int)
	go worker(c, id)
	return c
}

// worker
func channelDome1()  {
	var channels [10]chan<- int
	for i:=0; i<10 ; i++ {
		channels[i] = createWorker(i)
		/*channels[i] = make(chan int)
		go worker(i, channels[i])*/
	}
	for i:=0; i<10; i++ {
		channels[i] <- 'a' + i
	}

	time.Sleep(time.Millisecond)
}
func channelDome() {
	//var c chan int 	// 定义， c == nil
	c := make(chan int)
	go createWorker(0 )
	c <- 1	// 发送1
	c <- 2	// 发送2
	//n := <-c	// 将c中内容回收给n
	//println(n)
}

func worker(c chan int, id int) {
	for {
		n := <- c
		fmt.Printf("woker id %d, received %c\n", id, n)
	}
}

// send
func bufferedChannel()  {
	c := make(chan int, 3)
	go worker(c, 1)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	time.Sleep(time.Millisecond)
}

func channelClose()  {
	c := make(chan int, 3)
	go worker(c, 1)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	close(c)
	time.Sleep(time.Millisecond)
}
func main() {
	//channelDome() // 1 2
	//channelDome1()
	channelClose()
}
