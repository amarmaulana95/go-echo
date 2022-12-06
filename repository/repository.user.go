package repository

import (
	"github.com/amarmaulana95/go-echo/models"
	"github.com/amarmaulana95/go-echo/schemas"

	"gorm.io/gorm"
)

type repositoryUser struct {
	db *gorm.DB
}

// EntityCreate implements entity.EntityUsers
func (*repositoryUser) EntityCreate(input *schemas.SchemaUser) (*models.ModelUser, error) {
	panic("unimplemented")
}

func NewRepositoryUser(db *gorm.DB) *repositoryUser {
	return &repositoryUser{
		db: db,
	}
}

func (r *repositoryUser) EntityResults() (*[]models.ModelUser, error) {
	var users []models.ModelUser
	db := r.db.Model(&users)
	err := db.Debug().Find(&users).Error

	// if checkUser.RowsAffected < 1 {
	// 	return &users, checkUser.Error
	// }

	if err != nil {
		return &users, err
	}
	return &users, nil
}

func (r *repositoryUser) EntityResult(input *schemas.SchemaUser) (*models.ModelUser, error) {
	var user models.ModelUser
	user.ID = input.ID

	db := r.db.Model(&user)

	checkUser := db.Debug().First(&user, "id=?", input.ID)

	if checkUser.RowsAffected < 1 {
		return &user, checkUser.Error
	}

	return &user, nil
}
