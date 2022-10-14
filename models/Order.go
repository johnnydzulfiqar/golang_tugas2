package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model	
	CustomerName string `json:"customer_name" gorm:"not null;type:varchar(20)"`
	OrderedAt time.Time `json:"ordered_at" gorm:"not null"`
	Items []Item `gorm:"foreignKey:order_refer;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}