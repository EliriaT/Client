package client_elem

import "time"

const (
	OrderManagerUrl = "http://food-ordering:8084/"

	TimeUnit = time.Duration(float64(time.Millisecond) * 25)
)

var (
	ClientList       []Client
	NrClients        = 5
	NotifClientManag = make(chan int)
)
