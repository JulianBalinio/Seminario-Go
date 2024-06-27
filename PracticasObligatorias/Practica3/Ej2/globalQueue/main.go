package main

import (
	"fmt"
	"sync"
	"time"
	"Ej2/cash"
	"Ej2/client"
)

func main() {
	start := time.Now()
	cashQuant := 3
	clientsQuant := 10

	queue := make(chan *client.Client, clientsQuant) // Channel for clients

	var wg sync.WaitGroup

	// Create cashiers
	cashiers := make([]*cash.Cashier, cashQuant)
	for i := 0; i < cashQuant; i++ {
		cashiers[i] = &cash.Cashier{}
		cashiers[i].SetId(i + 1)
		
	}

	//atencion
	for _, c := range cashiers {
		go func(cashier *cash.Cashier) {
			for client := range queue {
				cashier.Serve(client, &wg)
				//wg.Done()
			}
		}(c)
	}

	//clientes a la queue
	for i := 1; i <= clientsQuant; i++ {
		client := &client.Client{}
		client.SetId(i)
		wg.Add(1)
		queue <- client
		time.Sleep(100 * time.Millisecond) //delay para dar tiempo a las goroutines
	}
	close(queue)

	wg.Wait()
	fmt.Println("Todos los clientes fueron atendidos")
	fmt.Printf("Tiempo de ejecucion: %s\n", time.Since(start))
}
