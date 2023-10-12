package models

type Product struct {
	ID       uint    `json:"id"`
	Name     string  `json:"name"`
	Category string  `json:"category"`
	Price    float64 `json:"price"`
}

type CartItem struct {
	ID        uint `json:"id"`
	ProductID uint `json:"product_id"`
	Quantity  uint `json:"quantity"`
}

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"`
}
