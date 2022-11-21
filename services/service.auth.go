package services

import (
	"github.com/amarmaulana95/go-echo/models"
	"github.com/amarmaulana95/go-echo/schemas"

	"github.com/amarmaulana95/go-echo/entity"
)

type serviceAuth struct {
	auth entity.EntityAuth
}

func NewServiceAuth(auth entity.EntityAuth) *serviceAuth {
	return &serviceAuth{auth: auth}
}

func (s *serviceAuth) EntityRegister(input *schemas.SchemaUser) (*models.ModelUser, error) {
	var schema schemas.SchemaUser

	// schema.FirstName = input.FirstName
	// schema.LastName = input.LastName
	schema.Username = input.Username
	schema.Email = input.Email
	schema.Password = input.Password
	// schema.Role = input.Role

	res, err := s.auth.EntityRegister(&schema)

	return res, err

}

func (s *serviceAuth) EntityLogin(input *schemas.SchemaUser) (*models.ModelUser, error) {
	var schema schemas.SchemaUser

	schema.Email = input.Email
	schema.Password = input.Password

	res, err := s.auth.EntityLogin(&schema)

	return res, err
}
