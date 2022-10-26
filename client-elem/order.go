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

type CookedOrder struct {
	OrderId        int              `json:"order_id"`
	IsReady        bool             `json:"is_ready"`
	EstimatedTime  int              `json:"estimated_waiting_time"`
	Priority       int              `json:"priority"`
	MaxWait        float32          `json:"max_wait"`
	CreatedTime    time.Time        `json:"created_time"`
	RegisteredTime time.Time        `json:"registered_time"`
	PreparedTime   time.Time        `json:"prepared_time"`
	CookingTime    time.Duration    `json:"cooking_time"`
	CookingDetails []kitchenFoodInf `json:"cooking_details"`
}
type kitchenFoodInf struct {
	FoodId int `json:"food_id"`
	CookId int `json:"cook_id"`
}

type raitingsResponses struct {
	ClientId int           `json:"client_id"`
	OrderId  int           `json:"order_id"`
	Orders   []orderRating `json:"orders"`
}

type orderRating struct {
	RestaurantId  int `json:"restaurant_id"`
	OrderId       int `json:"order_id"`
	Rating        int `json:"rating"`
	EstimatedTime int `json:"estimated_waiting_time"`
	WaitedTime    int `json:"waiting_time"`
}
