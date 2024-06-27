package main

import (
	"fmt"
)

func pxng(m chan string, str string, done chan bool) {
	m <- str
	done <- true
}

func main() {
	messages := make(chan string)
	done := make(chan bool)

	go func(){
		for i := 0; i < 5; i++ {
			go pxng(messages, "PING", done)
			<- done
			go pxng(messages, "PONG", done)
			<- done
		}
		close(messages)
	}()		
	
	for i := 0; i < 10; i++ {
		fmt.Println(<-messages)
	}
}