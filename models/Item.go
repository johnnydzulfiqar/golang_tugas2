package models

import (
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model	
	ItemCode string `json:"item_code" gorm:"not null;type:varchar(20)"`
	Description string `json:"description" gorm:"not null;type:varchar(20)"`
	Quantity int `json:"quantity" gorm:"not null;type:integer"`	
	OrderRefer int `json:"order_refer"`
}