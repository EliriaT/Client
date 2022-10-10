package main

import (
	client_elem "github.com/EliriaT/Client/client-elem"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)

	client_elem.InitClients()

	for i := range client_elem.ClientList {
		go client_elem.ClientList[i].OrderOnline()
	}

	wg.Wait()
}
