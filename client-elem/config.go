package client_elem

import "time"

const (
	OrderManagerUrl = "http://localhost:8084/"

	TimeUnit = time.Duration(float64(time.Millisecond) * 75)
)

var (
	ClientList []Client
	NrClients  = 6
)
