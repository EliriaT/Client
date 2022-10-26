package client_elem

import (
	"log"
	"sync/atomic"
	"time"
)

var ratingSum int32 = 0
var orderSum int32 = 0

func giveOrderStars(serveTime time.Duration, maxWait float64) int {
	//serveTimeMillisec := float64(serveTime)*1000 //time in milliseconds
	serveTimeNonUnit := float64(serveTime) / float64(TimeUnit)

	switch {
	case serveTimeNonUnit < maxWait:
		return 5
	case serveTimeNonUnit < maxWait*1.1:
		return 4
	case serveTimeNonUnit < maxWait*1.2:
		return 3
	case serveTimeNonUnit < maxWait*1.3:
		return 2
	case serveTimeNonUnit < maxWait*1.4:
		return 1
	default:
		return 0
	}
}

func CalculateRating(cookedOrder CookedOrder) (raiting int, waitingTime int) {
	atomic.AddInt32(&orderSum, 1)
	serveTime := time.Since(cookedOrder.RegisteredTime)
	waitingTime = int(serveTime) / int(TimeUnit)

	raiting = giveOrderStars(serveTime, float64(cookedOrder.MaxWait))
	atomic.AddInt32(&ratingSum, int32(raiting))

	log.Printf("--------Rating of ONLINE order %d is %d", cookedOrder.OrderId, raiting)
	log.Println("--------Average rating of clients is", float32(ratingSum)/float32(orderSum))
	return
}
