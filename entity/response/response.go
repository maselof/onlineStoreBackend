package response

type JSONResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type OrdersResponse struct {
	ID       int `json:"id" gorm:"column:order_id"`
	Products int `json:"product_id" gorm:"column:product_id"`
	Amount   int `amount:"amount" gorm:"column:amount"`
}

type UsersResponse struct {
	ID   int    `json:"id" gorm:"column:id"`
	Name string `json:"name" gorm:"column:name"`
}

type ProductsResponse struct {
	ID    int     `json:"id" gorm:"column:id"`
	Name  string  `json:"name" gorm:"column:name"`
	Price float64 `json:"price" gorm:"column:price"`
}
