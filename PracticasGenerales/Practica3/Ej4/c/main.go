package main

import (
	"fmt"
	"sync"
)

func send(messages chan string, str string, wg *sync.WaitGroup) {
	defer wg.Done()
	messages <- str
}

func main() {
	messages := make(chan string)
	var wg sync.WaitGroup

	go func() {
		for i := 0; i < 5; i++ {
			wg.Add(1)
			go send(messages, "PING", &wg)
			wg.Wait()

			wg.Add(1)
			go send(messages, "PONG", &wg)
			wg.Wait()
		}
		close(messages)
	}()

	for msg := range messages {
		fmt.Println(msg)
	}
}