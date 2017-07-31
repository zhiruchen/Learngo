package goroutines

import (
	"fmt"
)

func CloseRoutine() {
	dataChan := make(chan int, 5)
	c1 := make(chan struct{}, 1)
	c2 := make(chan struct{}, 2) // 为了不让主线程过早结束

	go func() {
		<-c1
		for {
			if elem, ok := <-dataChan; ok {
				fmt.Println("接受到: ", elem, "[receiver]")
			} else {
				break
			}
		}
		c2 <- struct{}{}
	}()

	go func() {
		for i := 1; i <= 5; i++ {
			dataChan <- i
			fmt.Println("发送: ", i)
		}

		close(dataChan)
		c1 <- struct{}{} // 关闭了通道，也能从中接受到值
		fmt.Println("Done")
		c2 <- struct{}{}
	}()

	// 只有上面两个线程都往c2发送了值，主线程才能接受到值，从而结束运行
	<-c2
	<-c2
}
