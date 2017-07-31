package goroutines

import (
	"fmt"
	"time"
)

var dc = make(chan string, 3)

func ForSelect() {
	c1 := make(chan struct{}, 1)
	c2 := make(chan struct{}, 2)
	go receive(dc, c1, c2)
	go send(dc, c1, c2)
	<-c2
	<-c2
}

func receive(c <-chan string, c1 <-chan struct{}, c2 chan<- struct{}) {
	<-c1

	time.Sleep(time.Second)
	for e := range c {
		fmt.Println("接受到: ", e)
	}
	c2 <- struct{}{}
}

func send(c chan<- string, c1, c2 chan<- struct{}) {
	for _, e := range []string{"a", "B", "C", "D"} {
		c <- e
		fmt.Println("发送: ", e)
		if e == "C" {
			c1 <- struct{}{}
			fmt.Println("发送一个同步信号")
		}
	}
	fmt.Println()
	time.Sleep(2 * time.Second)
	close(c)
	c2 <- struct{}{}
}
