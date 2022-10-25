package client_elem

type RestaurantsList struct {
	RestaurantsNum  int              `json:"restaurants"`
	RestaurantsData []RestaurantInfo `json:"restaurants_data"`
}

type RestaurantInfo struct {
	Id        int     `json:"restaurant_id"`
	Name      string  `json:"name"`
	MenuItems int     `json:"menu_items"`
	Menu      []Food  `json:"menu"`
	Rating    float32 `json:"rating"`
}

type Food struct {
	Id               int    `json:"id"`
	Name             string `json:"name"`
	PreparationTime  int    `json:"preparation-time"`
	Complexity       int    `json:"complexity"`
	CookingApparatus string `json:"cooking-apparatus"`
}
