package entity

import (
	"github.com/amarmaulana95/go-echo/models"
	"github.com/amarmaulana95/go-echo/schemas"
)

type EntityAuth interface {
	EntityRegister(input *schemas.SchemaUser) (*models.ModelUser, error)
	EntityLogin(input *schemas.SchemaUser) (*models.ModelUser, error)
}
