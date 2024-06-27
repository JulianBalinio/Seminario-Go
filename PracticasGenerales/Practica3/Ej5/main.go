package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const LIMIT = 3

func Producer(id int, ch chan <-int, wg *sync.WaitGroup){
	defer wg.Done()
	for i:= 0; i < LIMIT; i++ {
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		num := rand.Intn(101)
		fmt.Printf("Productor ID(%d) produjo el numero: %d. \n", id, num)
		ch <- num
	}
}

func Consumer(id int, ch <- chan int, wg *sync.WaitGroup){
	defer wg.Done()
	for i:= 0; i < LIMIT; i++ {
		num := <-ch
		fmt.Printf("Consumidor ID(%d) consumio el numero: %d.\n", id, num)
	}
}

func main() {
	ch := make(chan int)

	var wg sync.WaitGroup
	
	wg.Add(2)
	go Producer(1, ch, &wg)
	go Producer(2, ch, &wg)

	wg.Add(2)
	go Consumer(1, ch, &wg)
	go Consumer(2, ch, &wg)

	wg.Wait()
	close(ch)
}