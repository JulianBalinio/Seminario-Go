package cash

import (
	"fmt"
	"sync"
	"math/rand"
	"time"
	"Ej2/client"
)

type Cashier struct {
	id int
}

func (c *Cashier) SetId(n int) {
	c.id = n
}

func (c *Cashier) ID() int{
	return c.id
}

func (c *Cashier) Serve(client *client.Client, wg *sync.WaitGroup) {
	//defer wg.Done()
	fmt.Printf("Caja %d atendiendo a cliente %d\n", c.ID(), client.ID())
	time.Sleep(time.Duration(rand.Intn(2)) * time.Second)
	fmt.Printf("Caja %d termino de atender al cliente %d\n", c.ID(), client.ID())
	wg.Done()
}

