package client_elem

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type Client struct {
	Id              int `json:"id"`
	restaurantsMenu RestaurantsList
	generatedOrder  ClientOrderGenerated
	responseOrder   ClientOrderResponse
	ratingChan      chan orderRating
	orderWG         *sync.WaitGroup
	ratingResponse  raitingsResponses
}

func (c *Client) RequestMenu() {
	var restaurantsMenu RestaurantsList

	resp, err := http.Get(OrderManagerUrl + "menu")
	if err != nil {
		log.Fatal("Menu Request Failed: %s", err.Error())
		return
	}
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal("Menu Request Failed: %s", err.Error())
		return
	}
	json.Unmarshal(body, &restaurantsMenu)
	c.restaurantsMenu = restaurantsMenu
	log.Printf("Menu requested succesfully by client %d \n", c.Id)
}

func (c *Client) GenerateOrder() {
	var clientOrder ClientOrderGenerated
	clientOrder.ClientId = c.Id

	//generate random restaurants number
	nrRestaurants := c.restaurantsMenu.RestaurantsNum
	randomNrRestaurants := rand.Intn(nrRestaurants) + 1
	clientOrder.Orders = make([]OrderGenerated, randomNrRestaurants)

	//Example slice [4,1,3,2,0]
	randRestaurantsList := rand.Perm(nrRestaurants)
	randRestaurantsList = randRestaurantsList[0:randomNrRestaurants]

	for i, restaurantId := range randRestaurantsList {
		var restaurantData RestaurantInfo
		//find the restaurant data with that specific ID
		for _, restaurInfo := range c.restaurantsMenu.RestaurantsData {
			if restaurInfo.Id == restaurantId+1 {
				restaurantData = restaurInfo
			}
		}
		restaurantOrder := GenerateOneOrder(restaurantData, restaurantId+1)
		clientOrder.Orders[i] = restaurantOrder
	}

	c.generatedOrder = clientOrder
	log.Printf("Client %d generated order %v", c.Id, c.generatedOrder)
}

func (c *Client) SendOrder() {
	var clientOrderResponse ClientOrderResponse
	reqBody, err := json.Marshal(c.generatedOrder)
	if err != nil {
		log.Printf(err.Error())
		return
	}

	resp, err := http.Post(OrderManagerUrl+"order", "application/json", bytes.NewBuffer(reqBody))

	//if response failed then do not wait and do not pick up. I should return an err and check everytime for that error
	if err != nil {
		log.Printf("Request Failed: %s", err.Error())
		return
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Printf("Reading Response reading Failed: %s", err.Error())
		return
	}
	_ = json.Unmarshal(body, &clientOrderResponse)
	c.responseOrder = clientOrderResponse

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf(err.Error())
			return
		}
	}(resp.Body)

	log.Printf("The order of the client %d was sent to Orderds Manager .", c.Id)
	log.Printf("Got response: %+v", clientOrderResponse)
}

func (c *Client) WaitForOrders() {
	c.ratingChan = make(chan orderRating, len(c.responseOrder.Orders))
	c.orderWG = new(sync.WaitGroup)
	c.orderWG.Add(len(c.responseOrder.Orders))

	for _, order := range c.responseOrder.Orders {
		log.Printf(" ------- Client %d is waiting %d timeunits for order %d -------", c.Id, order.EstimatedWaitingTime, order.OrderId)
		go c.PickUpOnlineOrder(order)
	}
	c.orderWG.Wait()

	close(c.ratingChan)
	for orderWithRating := range c.ratingChan {
		c.ratingResponse.Orders = append(c.ratingResponse.Orders, orderWithRating)
	}
	c.ratingResponse.ClientId = c.Id
	c.ratingResponse.OrderId = int(c.responseOrder.OrderId)

}

func (c *Client) PickUpOnlineOrder(order OrderResponse) {
	var cookedOrderResponse CookedOrder

	strId := strconv.Itoa(order.OrderId)
	if strId == "" {
		log.Fatal("Error, can not convert order id of to string")
	}
	//Wait for the estimated waiting time
	time.Sleep(TimeUnit * time.Duration(order.EstimatedWaitingTime))

	resp, err := http.Get(order.RestaurantAddress + "v2/order/" + strId)
	if err != nil {
		log.Fatal(" Request for the Online Cooked Order Failed: %s", err.Error())
		return
	}
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal("Reading the response of the online picked order failed : %s", err.Error())
		return
	}
	json.Unmarshal(body, &cookedOrderResponse)

	if cookedOrderResponse.IsReady == true {
		log.Printf("----------------Client %d succesfully picked online order %d %+v\n", c.Id, order.OrderId, cookedOrderResponse)
		rating, waitedTime := CalculateRating(cookedOrderResponse)
		orderRatingResponse := orderRating{
			RestaurantId:  order.RestaurantId,
			OrderId:       order.OrderId,
			Rating:        rating,
			EstimatedTime: order.EstimatedWaitingTime,
			WaitedTime:    waitedTime,
		}
		c.ratingChan <- orderRatingResponse
		c.orderWG.Done()
	} else {

		go c.PickUpOnlineOrder(order)
	}

}

func (c *Client) SendRaitings() {
	reqBody, err := json.Marshal(c.ratingResponse)
	if err != nil {
		log.Printf(err.Error())
		return
	}
	_, err = http.Post(OrderManagerUrl+"rating", "application/json", bytes.NewBuffer(reqBody))

	if err != nil {
		log.Printf("Raiting Post Request Failed: %s", err.Error())
		return
	}
}

func (c *Client) OrderOnline() {

	time.Sleep(TimeUnit * time.Duration(rand.Intn(500)+30))
	c.RequestMenu()
	c.GenerateOrder()
	c.SendOrder()
	c.WaitForOrders()
	c.SendRaitings()
	//Notifying that a Client finished his job and another client can be generated
	NotifClientManag <- c.Id

}
