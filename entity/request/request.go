package request

type OrderRequest struct {
	ID     int              `json:"id"`
	UserID int              `json:"user_id"`
	Items  []ProductRequest `json:"items"`
}

type ProductRequest struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type UserRequest struct {
	Name string `json:"name"`
}

type CastRequest struct {
	UserID     int   `json:"user_id"`
	ProductsID []int `json:"products_id"`
}
