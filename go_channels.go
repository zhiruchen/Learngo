package main

import (
	"fmt"
	"time"
)

func pinger(c chan string) {
	for i := 0; ; i++ {
		c <- "ping" // <- send "ping" to c
	}
}

func ponger(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

func printer(c chan string) {
	for {
		msg := <-c // receive a msg from c and store in msg
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	// var c chan string = make(chan string)
	// go pinger(c)
	// go ponger(c)
	// go printer(c)

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "from msg1"
			time.Sleep(time.Second * 2)
		}
	}()

	go func() {
		for {
			c2 <- "from msg2"
			time.Sleep(time.Second * 2)
		}
	}()

	// select 选择第一个准备好的channel，从channel中获取msg
	// 如果两个都准备好了，则随机选择一个来获取
	// 如果都没有准备好，下面表达式就阻塞直到有一个可用
	go func() {
		for {
			select {
			case msg1 := <-c1:
				fmt.Println(msg1)
			case msg2 := <-c2:
				fmt.Println(msg2)
			}
		}
	}()

	var input string
	fmt.Scanln(&input)
}
