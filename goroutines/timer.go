package goroutines

import (
	"fmt"
	"time"
)

func TimeExpiration() {
	timer := time.NewTimer(2 * time.Second)
	fmt.Println("present time: ", time.Now())

	expirationTime := <-timer.C // 阻塞直到定时器到期
	fmt.Printf("ExpirationTime: %v\n", expirationTime)
	fmt.Printf("Stop timer, %v\n", timer.Stop())
}

func ChannelTimeout() {
	c := make(chan int, 1)

	go func() {
		time.Sleep(time.Second)
		c <- 1
	}()

	select {
	case e := <-c:
		fmt.Printf("Received: %v\n", e)
	case <-time.After(500 * time.Millisecond):
		fmt.Println("time out")
	}
}
