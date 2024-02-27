package dbmodel

type Products struct {
	ID    int    `gorm:"column:id"`
	Name  string `gorm:"column:name"`
	Price string `gorm:"column:price"`
}

type Order struct {
	ID     int `gorm:"column:id"`
	UserID int `gorm:"column:user_id"`
}
