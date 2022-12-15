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

func (s *serviceUser) EntityResultID(input *schemas.SchemaUser) (*models.ModelUser, error) {
	var user schemas.SchemaUser
	user.ID = input.ID

	res, err := s.users.EntityResultID(&user)

	return res, err
}

func (s *serviceUser) EntityResultIDSearch(search string) (*[]models.ModelUser, error) {
	res, err := s.users.EntityResultIDSearch(search)
	return res, err

	// for i, _ := range arrayProducts {
	// 	VarProducts.ID = arrayProducts[i].ID
	// 	VarProducts.Username = arrayProducts[i].Username
	// 	VarProducts.Email = arrayProducts[i].Email

	// 	responseProducts = schemas.SchemaUser{
	// 		VarProducts.ID,
	// 		VarProducts.Username,
	// 		VarProducts.Email
	// 	} //
	// 	arrResponAnalisa = append(arrResponAnalisa, responseProducts)
	// }
}

func (s *serviceUser) EntityResultAll(search string, limit uint64, offset uint64) (*[]models.ModelUser, error) {
	res, err := s.users.EntityResultAll(search, limit, offset)
	return res, err
}

// func (s *serviceUser) EntityResultAllTotal(search string) (*models.ModelUser, error) {
// 	res, err := s.users.EntityResultAllTotal(search)
// 	return res, err
// }
