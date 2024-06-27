package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan string)

	go func() {
		for i := 0; i < 5; i++ {
			messages <- "PING"
			time.Sleep(100 * time.Millisecond)
			messages <- "PONG"
			time.Sleep(100 * time.Millisecond)
		}
		close(messages)
	}()

	for msg := range messages {
		fmt.Println(msg)
	}
}