package goroutines

import (
	"fmt"
	"math"
	"time"
)

func putSthToChannel(sth interface{}, c chan interface{}) {
	c <- sth
}

func calcPrime(nl []int, count int, done chan bool) {
Loop:
	for i := 2; i <= count; i++ {
		var j int
		for j = 2; float64(j) <= math.Sqrt(float64(i)); j++ {
			if i%j == 0 {
				continue Loop
			}
		}
		nl = append(nl, i)
	}
	done <- true
}

func calcuPrime2(count int, nl chan int) {
Loop:
	for i := 2; i <= count; i++ {
		var j int
		for j = 2; float64(j) <= math.Sqrt(float64(i)); j++ {
			if i%j == 0 {
				continue Loop
			}
		}
		nl <- i
	}
	close(nl)
}

func CallPut() {
	c := make(chan interface{})
	go putSthToChannel("channel", c)
	v := <-c
	fmt.Println(v)

	go putSthToChannel(1000000, c)
	v = <-c
	fmt.Println(v)

	go putSthToChannel(struct{ a string }{a: "1"}, c)
	v = <-c
	fmt.Println(v)

	go putSthToChannel(time.Now().Unix(), c)
	v = <-c
	fmt.Println(v)
}

func CalcuPrime() {
	nl := []int{}
	c := make(chan bool)
	go calcPrime(nl, 100, c)

	v := <-c
	fmt.Println(v)
	fmt.Println(nl)

	for _, n := range nl {
		fmt.Println(n)
	}
}

func CalcuPrime2(count int) {
	// count := 10
	nl := make(chan int, count)
	go calcuPrime2(count, nl)
	for x := range nl {
		fmt.Println(x)
	}
}
