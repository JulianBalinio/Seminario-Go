package main

import (
    "fmt"
)

func main() {
    fmt.Println("Inicia Goroutine del main")
	done := make(chan bool)
    go hello(done)
    <- done
	fmt.Println("Termina Goroutine del main")
	
}

func hello(done chan bool) {
    fmt.Println("Inicia Goroutine de hello")
    for i := 0; i < 3; i++ {
        fmt.Println(" Hello world")
    }
	
    fmt.Println("Termina Goroutine de hello")
	done <- true
}