package client_elem

import "time"

type ClientOrderGenerated struct {
	ClientId int              `json:"client_id"`
	Orders   []OrderGenerated `json:"orders"`
}

type ClientOrderResponse struct {
	OrderId int32           `json:"order_id"`
	Orders  []OrderResponse `json:"orders"`
}

type OrderGenerated struct {
	RestaurantId int       `json:"restaurant_id"`
	Items        []int     `json:"items"`
	Priority     int       `json:"priority"`
	MaxWait      float32   `json:"max_wait"`
	CreatedTime  time.Time `json:"created_time"`
}

type OrderResponse struct {
	RestaurantId         int       `json:"restaurant_id"`
	RestaurantAddress    string    `json:"restaurant_address"`
	OrderId              int       `json:"order_id"`
	EstimatedWaitingTime int       `json:"estimated_waiting_time"`
	CreatedTime          time.Time `json:"created_time"`
	RegisteredTime       time.Time `json:"registered_time"`
}
