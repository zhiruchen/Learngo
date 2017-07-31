package goroutines

import (
	"fmt"
	"time"
)

var strChan = make(chan string, 3)

func SyncRoutine() {
	c1 := make(chan struct{}, 1)
	c2 := make(chan struct{}, 2) // 为了不让主线程过早结束

	go func() {
		<-c1

		fmt.Println("因c1没有足够的元素, 接受操作阻塞了")
		time.Sleep(time.Second)

		for {
			if elem, ok := <-strChan; ok {
				fmt.Println("接受到: ", elem, "[receiver]")
			} else {
				break
			}
		}
		fmt.Println("接受者线停止")
		c2 <- struct{}{}
	}()

	go func() {
		for _, e := range []string{"1", "2", "3", "c"} {
			strChan <- e
			fmt.Println("发送: ", e)

			if e == "3" {
				c1 <- struct{}{} // 通知接受者线程接受发送的元素
				fmt.Println("发送了一个同步信号")
			}
		}
		fmt.Println("暂停2秒")
		time.Sleep(time.Second * 2)
		close(strChan)
		c2 <- struct{}{}
	}()

	// 只有上面两个线程都往c2发送了值，主线程才能接受到值，从而结束运行
	<-c2
	<-c2
}
