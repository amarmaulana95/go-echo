package models

import (
	"time"
)

type ModelProduct struct {
	ID          string    `json:"id" gorm:"primary_key"`
	Name        string    `json:"name" gorm:"type:varchar;"`
	Price       string    `json:"price" gorm:"type:varchar;"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (ModelProduct) TableName() string {
	return "products"
}
