package entity

import (
	"github.com/amarmaulana95/go-echo/models"
)

type EntityProducts interface {
	EntityResults() (*[]models.ModelProduct, error)
	EntityResultAll(search string, limit uint64, offset uint64) (*[]models.ModelProduct, error)
	EntityResultAllTotal(search string) uint64
}
