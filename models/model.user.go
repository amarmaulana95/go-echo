package models

import (
	"time"

	"github.com/amarmaulana95/go-echo/pkg"
	"gorm.io/gorm"
)

type ModelUser struct {
	ID string `json:"id" gorm:"primary_key"`
	// FirstName string    `json:"first_name" gorm:"type:varchar;  not null"`
	// LastName  string    `json:"last_name" gorm:"type:varchar; not null"`
	Username string `json:"username" gorm:"type:varchar;"`
	Email    string `json:"email" gorm:"type:varchar; unique; not null"`
	Password string `json:"password" gorm:"type:varchar; not null"`
	// Role      string    `json:"role" gorm:"type:varchar; not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (ModelUser) TableName() string {
	return "users"
}

func (m *ModelUser) BeforeCreate(db *gorm.DB) error {
	// m.ID = uuid.NewString()
	m.Password = pkg.HashPassword(m.Password)
	m.CreatedAt = time.Now()
	return nil
}

func (m *ModelUser) BeforeUpdate(db *gorm.DB) error {
	// m.ID = uuid.NewString()
	m.Password = pkg.HashPassword(m.Password)
	m.UpdatedAt = time.Now()
	return nil
}
