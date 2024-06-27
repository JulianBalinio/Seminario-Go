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
	cashQuant := 10
	clientsQuant := 50

	// Crear canales individuales para cada caja
	cashierQueues := make([]chan *client.Client, cashQuant)
	for i := 0; i < cashQuant; i++ {
		cashierQueues[i] = make(chan *client.Client, clientsQuant)
	}

	var wg sync.WaitGroup

	// Crear cajeros
	cashiers := make([]*cash.Cashier, cashQuant)
	for i := 0; i < cashQuant; i++ {
		cashiers[i] = &cash.Cashier{}
		cashiers[i].SetId(i + 1)

		go func(cashier *cash.Cashier, queue chan *client.Client) {
			for client := range queue {
				cashier.Serve(client, &wg)
				//wg.Done()
			}
		}(cashiers[i], cashierQueues[i])
	}

	//Asignar clientes(rr)
	for i := 1; i <= clientsQuant; i++ {
		client := &client.Client{}
		client.SetId(i)
		wg.Add(1)
		cashierQueues[(i-1)%cashQuant] <- client
		time.Sleep(100 * time.Millisecond) //delay para dar tiempo a las goroutines
	}

	for i := 0; i < cashQuant; i++ {
		close(cashierQueues[i])
	}

	wg.Wait()
	fmt.Println("Todos los clientes fueron atendidos")
	fmt.Printf("Tiempo de ejecución: %s\n", time.Since(start))
}
