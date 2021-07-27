package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Print("trabalhando...")
	time.Sleep(time.Second)
	fmt.Println("finalizado")

	done <- true
}

func main() {

	done := make(chan bool, 1)
	go worker(done)

	<-done
}