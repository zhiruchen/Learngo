package goroutines

import "fmt"

var c1 chan int
var c2 chan int
var chanList = []chan int{c1, c2}
var numbers = []int{1, 2, 3, 4, 5}

func SelectEval() {
	select {
	default:
		fmt.Println("Default")
	case getChan(0) <- getNumber(0):
		fmt.Println("1th case selected")
	case getChan(1) <- getNumber(1):
		fmt.Println("2nd case is selected")
	}
}

func getNumber(i int) int {
	fmt.Printf("numbers[%d]\n", i)
	return numbers[i]
}

func getChan(i int) chan int {
	fmt.Printf("channels[%d]\n", i)
	return chanList[i]
}

func SelectRandom() {
	chanCap := 5
	intChan := make(chan int, chanCap)

	for i := 0; i < chanCap; i++ {
		select {
		case intChan <- 1:
		case intChan <- 2:
		case intChan <- 3:
		}
	}

	for i := 0; i < chanCap; i++ {
		fmt.Printf("%d\n", <-intChan)
	}
}

func SelectFor() {
	intChan := make(chan int, 10)

	for i := 0; i < 10; i++ {
		intChan <- i
	}

	close(intChan)

	sc := make(chan struct{}, 1)
	go func() {
	Loop:
		for {
			select {
			case e, ok := <-intChan:
				if !ok {
					fmt.Println("end.")
					break Loop
				}
				fmt.Printf("Received: %v\n", e)
			}
		}
		sc <- struct{}{}
	}()
	<-sc
}
