package goroutines

import (
	"fmt"
	"time"
)

func SendMesageToChannel() {
	msgChan := make(chan string)

	go func() { msgChan <- "this is a message," }()

	msg := <-msgChan
	fmt.Println(msg)
}

func MoveBoxes(boxCounts int) {
	syncChan := make(chan int, boxCounts)

	for i := 1; i <= boxCounts; i++ {
		go moveBox(i)
		syncChan <- i
	}

	for i := 1; i <= boxCounts; i++ {
		<-syncChan
	}
	fmt.Println("ending move boxes")
}

func moveBox(i int) {
	fmt.Println("moving box", i)
}

func SyncChannel() {
	done := make(chan bool)

	go worker(done)

	<-done
}

func worker(done chan bool) {
	fmt.Println("starting working")
	time.Sleep(2 * time.Second)
	fmt.Println("ending working")

	done <- true
}

func SendJobs() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			job, more := <-jobs
			if more {
				fmt.Println("received job", job)
			} else {
				fmt.Println("there is no jobs")
				done <- true
				break
			}
		}
	}()

	for i := 1; i <= 3; i++ {
		jobs <- i
		fmt.Println("sending job ", i)
	}
	close(jobs)
	<-done
}

func MakeTicker() {
	ticker := time.NewTicker(time.Millisecond * 500)

	go func() {
		for t := range ticker.C {
			fmt.Println("tick at: ", t)
		}
	}()

	time.Sleep(time.Millisecond * 1600)
	ticker.Stop()

	fmt.Println("ticker stoped")
}

func WorkerPool() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for w := 1; w <= 3; w++ {
		go intWorker(w, jobs, results)
	}

	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	for x := 1; x <= 5; x++ {
		fmt.Println(<-results)
	}
}

func intWorker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Println("worker: ", id, " started job: ", job)
		time.Sleep(time.Second)
		results <- job * 2
	}
}
