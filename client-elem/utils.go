package client_elem

import (
	"math"
	"math/rand"
	"time"
)

func GenerateOneOrder(restaurantMenu RestaurantInfo, restaurantId int) OrderGenerated {
	var order OrderGenerated

	Foods := restaurantMenu.Menu
	nrTotalFoods := restaurantMenu.MenuItems
	nrItems := rand.Intn(nrTotalFoods) + 1

	for i := 0; i < 3; i++ {
		if nrItems > 5 {
			nrItems = rand.Intn(nrTotalFoods) + 1
		} else {
			break
		}
	}

	foodList := make([]int, nrItems)
	maxWait := 0

	for i := 0; i < nrItems; i++ {
		foodId := rand.Intn(nrTotalFoods) + 1
		//in the json the id starts from 1
		foodList[i] = foodId

		if prepTime := Foods[foodId-1].PreparationTime; prepTime > maxWait {
			maxWait = prepTime
		}

	}

	order.RestaurantId = restaurantId
	order.Items = foodList
	order.Priority = int(math.Round(float64(nrItems) / 2))
	order.MaxWait = float32(maxWait) * 1.8
	order.CreatedTime = time.Now()

	return order
}
