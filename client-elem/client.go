package client_elem

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type Client struct {
	Id              int `json:"id"`
	restaurantsMenu RestaurantsList
	generatedOrder  ClientOrderGenerated
	responseOrder   ClientOrderResponse
}

func (c *Client) RequestMenu() {
	var restaurantsMenu RestaurantsList

	resp, err := http.Get(OrderManagerUrl + "menu")
	if err != nil {
		log.Printf("Menu Request Failed: %s", err.Error())
		return
	}
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Printf("Menu Request Failed: %s", err.Error())
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

	//Example slice [4,1,3,2]
	randRestaurantsList := rand.Perm(nrRestaurants)
	randRestaurantsList = randRestaurantsList[0:randomNrRestaurants]

	for i, restaurantId := range randRestaurantsList {
		restaurantOrder := GenerateOneOrder(c.restaurantsMenu.RestaurantsData[restaurantId], restaurantId+1)
		clientOrder.Orders[i] = restaurantOrder
	}
	log.Println(clientOrder.Orders)
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

	if err != nil {
		log.Printf("Request Failed: %s", err.Error())
		return
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Printf("Menu Response reading Failed: %s", err.Error())
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
	log.Printf("Got response: %v", clientOrderResponse)
}

func (c *Client) WaitForOrders() {
	for _, order := range c.responseOrder.Orders {
		log.Println("---Waiting ", order.EstimatedWaitingTime)
		go time.Sleep(TimeUnit * time.Duration(order.EstimatedWaitingTime) * 3)
	}
	log.Printf("Client %d finished waiting for the orders %v", c.Id, c.responseOrder.Orders)
}

func (c *Client) OrderOnline() {
	for {
		time.Sleep(TimeUnit * time.Duration(rand.Intn(500)+30))
		c.RequestMenu()
		c.GenerateOrder()
		c.SendOrder()
		c.WaitForOrders()

	}

}
