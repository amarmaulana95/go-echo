package entity

import (
	"github.com/amarmaulana95/go-echo/models"
	"github.com/amarmaulana95/go-echo/schemas"
)

type EntityUsers interface {
	EntityResults() (*[]models.ModelUser, error)
	EntityResultID(input *schemas.SchemaUser) (*models.ModelUser, error)
	EntityResultIDSearch(search string) (*[]models.ModelUser, error)
	EntityResultAll(search string, limit uint64, offset uint64) (*[]models.ModelUser, error)
	EntityResultAllTotal(search string) uint64
}
