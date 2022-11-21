package entity

import (
	"github.com/amarmaulana95/go-echo/models"
	"github.com/amarmaulana95/go-echo/schemas"
)

type EntityUsers interface {
	EntityResults() (*[]models.ModelUser, error)
	EntityCreate(input *schemas.SchemaUser) (*models.ModelUser, error)
	EntityResult(input *schemas.SchemaUser) (*models.ModelUser, error)
	EntityUpdate(input *schemas.SchemaUser) (*models.ModelUser, error)
	EntityDelete(input *schemas.SchemaUser) (*models.ModelUser, error)
}
