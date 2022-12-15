package entity

import (
	"github.com/amarmaulana95/go-echo/models"
)

type EntityProducts interface {
	EntityResults() (*[]models.ModelProduct, error)
	// FindByID(search string, ID string) (*[]models.ModelProduct, error)
}
