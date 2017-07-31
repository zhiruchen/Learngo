package goroutines

import (
	"fmt"
	"runtime"
	"sync"
)

// RunGoRoutine run goroutine
func RunGoRoutine() {
	runtime.GOMAXPROCS(1)

	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start goroutine")

	go func() {
		defer wg.Done()

		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	go func() {
		defer wg.Done()

		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	fmt.Println("waiting to finish...")
	wg.Wait()

	fmt.Println("\n Terminating Program")
}

// TimeSliceGoroutine goroutine调度器在不同goroutine之间的切换
func TimeSliceGoroutine() {
	runtime.GOMAXPROCS(1)

	var wg sync.WaitGroup
	wg.Add(2)

	printPrime := func(prefix string) {
		defer wg.Done()

	next:
		for outter := 1; outter < 5000; outter++ {
			for inner := 2; inner < outter; inner++ {
				if outter%inner == 0 {
					continue next
				}
			}
			fmt.Printf("%s:%d\n", prefix, outter)
		}
		fmt.Println("Completed ", prefix)
	}

	go printPrime("A")
	go printPrime("B")

	fmt.Println("wait for the  goroutine to finish")
	wg.Wait()

	fmt.Println("Terminating Program...")
}
