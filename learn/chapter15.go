package main

import "fmt"

func main() {
	// 非缓冲通道
	ch := make(chan int)
	// 必须要有消费侧等待通道写入 才能向通道写入
	// 不然会报异常
	go recv(ch)
	ch <- 15
	fmt.Println("---------完毕")

	//缓冲通道
	ch2 := make(chan int, 1)
	ch2 <- 185
	//超出缓冲区异常 ch2<-2
	fmt.Println("通道2:", <-ch2)
}

// 接收
func recv(c chan int) {
	fmt.Println("接收： ", <-c)
}
