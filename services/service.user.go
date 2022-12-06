package services

import (
	"github.com/amarmaulana95/go-echo/entity"
	"github.com/amarmaulana95/go-echo/models"
	"github.com/amarmaulana95/go-echo/schemas"
)

type serviceUser struct {
	users entity.EntityUsers
}

func NewServiceUser(users entity.EntityUsers) *serviceUser {
	return &serviceUser{
		users: users,
	}
}

func (s *serviceUser) EntityResults() (*[]models.ModelUser, error) {
	res, err := s.users.EntityResults()
	if err != nil {
		return res, err
	}
	return res, err
}

func (s *serviceUser) EntityResult(input *schemas.SchemaUser) (*models.ModelUser, error) {
	var user schemas.SchemaUser
	user.ID = input.ID

	res, err := s.users.EntityResult(&user)

	return res, err
}
