package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func request(index int, ch chan string) {
	time.Sleep(time.Duration(index) * time.Second)
	s := fmt.Sprintf("编号%d完成", index)
	ch <- s
	defer wg.Done()
}
func main() {
	ch := make(chan string, 10)

	fmt.Println(ch, len(ch))
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go request(i, ch)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	for ret := range ch {
		fmt.Println("长度:", len(ch))
		fmt.Println("值：", ret)
	}

}
