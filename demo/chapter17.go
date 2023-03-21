package main

import (
	"fmt"
	"time"
)

func request2(index int, ch chan string) {
	time.Sleep(time.Duration(index) * time.Second)
	s := fmt.Sprintf("编号%d完成", index)
	ch <- s
}
func main() {
	ch := make(chan string, 1)
	for i := 0; i < 20; i++ {
		go request2(i, ch)
	}
	for {
		select {
		case i := <-ch:
			println(i)
		default:
			time.Sleep(time.Second)
			fmt.Println("无数据: ", time.Now())

		}
	}
}
