package model

import "time"

type Store struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string    `json:"name"`
	Logo      string    `json:"logo"`
	Theme     string    `json:"theme"`
	Banner    string    `json:"banner"`
	Notice    string    `json:"notice"`
	CreatedAt time.Time `json:"created_at"`
}

type Category struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string    `json:"name"`
	Icon      string    `json:"icon"`
	Sort      int       `json:"sort"`
	CreatedAt time.Time `json:"created_at"`
}

type Product struct {
	ID          int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string    `json:"name"`
	Price       float64   `json:"price"`
	Image       string    `json:"image"`
	CategoryID  int       `json:"category_id"`
	Description string    `json:"description"`
	Stock       int       `json:"stock"`
	Status      int       `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	Category    Category  `json:"category,omitempty" gorm:"foreignKey:CategoryID"`
}

type Cart struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID    string    `json:"user_id"`
	ProductID int       `json:"product_id"`
	Quantity  int       `json:"quantity"`
	CreatedAt time.Time `json:"created_at"`
	Product   Product   `json:"product,omitempty" gorm:"foreignKey:ProductID"`
}

type Order struct {
	ID           int         `json:"id" gorm:"primaryKey;autoIncrement"`
	OrderNo      string      `json:"order_no"`
	UserID       string      `json:"user_id"`
	TotalAmount  float64     `json:"total_amount"`
	Status       int         `json:"status"`
	ContactName  string      `json:"contact_name"`
	ContactPhone string      `json:"contact_phone"`
	ContactAddr  string      `json:"contact_addr"`
	Remark       string      `json:"remark"`
	CreatedAt    time.Time   `json:"created_at"`
	Items        []OrderItem `json:"items,omitempty" gorm:"foreignKey:OrderID"`
}

type OrderItem struct {
	ID          int     `json:"id" gorm:"primaryKey;autoIncrement"`
	OrderID     int     `json:"order_id"`
	ProductID   int     `json:"product_id"`
	ProductName string  `json:"product_name"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
}
