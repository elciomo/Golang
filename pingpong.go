package main

import (
	"fmt"
	"time"
)

func prcPing(done chan bool) {
	fmt.Println("Ping")
	time.Sleep(time.Second)

	done <- true
}

func prcPong(done chan bool) {
	fmt.Println("Pong")
	time.Sleep(time.Second)

	done <- true
}

func main() {
	for i := 0; ; i++ {
		done := make(chan bool, 1)

		go prcPing(done)
		<-done

		go prcPong(done)
		<-done
	}
}
