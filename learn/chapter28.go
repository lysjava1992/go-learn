package main

import (
	"fmt"
	"time"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
		time.Sleep(time.Second * 5)
	}
	close(c)

}
func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	//for {
	//	select {
	//	case i := <-c:
	//		fmt.Println("----------", i)
	//	default:
	//		// 如果两个通道都没有可用的数据，则执行这里的语句
	//		fmt.Println("no message received")
	//	}
	//}
	//for i := range c 会不断从信道接收值，直到它被关闭。
	for i := range c {
		fmt.Println(i)
	}
}
