package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("jobs received", j)
			} else {
				fmt.Println("all jobs received")
				done <- true
				return
			}
		}
	}()

	for i := 0; i < 5; i++ {
		jobs <- i
		fmt.Println("Sent job ", i)
	}

	close(jobs)
	fmt.Println("sent all jobs")

	<-done

}
