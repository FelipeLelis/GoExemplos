package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("job recebido", j)
			} else {
				fmt.Println("recebido todos os jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("enviando job", j)
	}
	close(jobs)
	fmt.Println("enviando todos os jobs")

	<-done
}