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

	return res, err
}

func (s *serviceUser) EntityCreate(input *schemas.SchemaUser) (*models.ModelUser, error) {
	var user schemas.SchemaUser

	// user.FirstName = input.FirstName
	// user.LastName = input.LastName
	user.Username = input.Username
	user.Email = input.Email
	user.Password = input.Password
	// user.Role = input.Role

	res, err := s.users.EntityCreate(&user)

	return res, err
}

func (s *serviceUser) EntityResult(input *schemas.SchemaUser) (*models.ModelUser, error) {
	var user schemas.SchemaUser
	user.ID = input.ID

	res, err := s.users.EntityResult(&user)

	return res, err
}

func (s *serviceUser) EntityUpdate(input *schemas.SchemaUser) (*models.ModelUser, error) {
	var user schemas.SchemaUser

	// user.FirstName = input.FirstName
	// user.LastName = input.LastName

	user.Email = input.Email
	user.Password = input.Password
	// user.Role = input.Role

	res, err := s.users.EntityUpdate(&user)

	return res, err
}

func (s *serviceUser) EntityDelete(input *schemas.SchemaUser) (*models.ModelUser, error) {
	var user schemas.SchemaUser
	user.ID = input.ID

	res, err := s.users.EntityDelete(&user)

	return res, err
}
