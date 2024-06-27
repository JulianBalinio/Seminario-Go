package main

import (
	"fmt"
	"math/rand"
	"sync"
)

const LIMIT = 5

func Insert(ch chan <- int, wg *sync.WaitGroup){
	defer wg.Done()
	for i:= 0; i < rand.Intn(6); i++ {
		ch <- rand.Intn(1001)
	}
}

func main(){
	var wg sync.WaitGroup

	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	ch1Quant := 0
	ch2Quant := 0
	ch3Quant := 0
	
	wg.Add(3)
	go Insert(ch1, &wg)
	go Insert(ch2, &wg)
	go Insert(ch3, &wg)
	
	
	go func() {
		wg.Wait()
		close(ch1)
		close(ch2)
		close(ch3)
	}()

	done := 0
	for done < 3 {
		select {
		case val, ok := <-ch1:
			if ok {
				ch1Quant++
				fmt.Printf("%d from ch1\n", val)
			} else {
				done++
				ch1 = nil
			}
		case val, ok := <-ch2:
			if ok {
				ch2Quant++
				fmt.Printf("%d from ch2\n", val)
			} else {
				done++
				ch2 = nil
			}
		case val, ok := <-ch3:
			if ok {
				ch3Quant++
				fmt.Printf("%d from ch3\n", val)
			} else {
				done++
				ch3 = nil
			}
		}
	}

	fmt.Printf("ch1Quant: %d\n", ch1Quant)
	fmt.Printf("ch2Quant: %d\n", ch2Quant)
	fmt.Printf("ch3Quant: %d\n", ch3Quant)

}