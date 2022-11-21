package repository

import (
	"github.com/amarmaulana95/go-echo/models"
	"github.com/amarmaulana95/go-echo/pkg"
	"github.com/amarmaulana95/go-echo/schemas"
	"golang.org/x/crypto/bcrypt"

	"gorm.io/gorm"
)

type repositoryAuth struct {
	db *gorm.DB
}

func NewRepositoryAuth(db *gorm.DB) *repositoryAuth {
	return &repositoryAuth{db: db}
}

func (r *repositoryAuth) EntityRegister(input *schemas.SchemaUser) (*models.ModelUser, error) {
	var user models.ModelUser
	// user.FirstName = input.FirstName
	// user.LastName = input.LastName
	user.Username = input.Username
	user.Email = input.Email
	user.Password = input.Password
	// user.Role = input.Role

	db := r.db.Model(&user)

	checkEmailExist := db.Debug().First(&user, "email=?", input.Email)
	// fmt.Println(checkEmailExist)

	if checkEmailExist.RowsAffected > 0 {
		return &models.ModelUser{}, nil
	}

	addUser := db.Debug().Create(&user).Commit()

	if addUser.RowsAffected < 1 {
		return &user, nil
	}

	return &user, nil
}

func (r *repositoryAuth) EntityLogin(input *schemas.SchemaUser) (*models.ModelUser, error) {
	var user models.ModelUser
	var err error
	user.Email = input.Email
	user.Password = input.Password

	db := r.db.Model(&user)
	err = db.Debug().First(&user, "email=?", input.Email).Error
	if err != nil {
		return &user, err
	}
	err = pkg.VerifyPassword(user.Password, input.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {

		return &user, err
	}
	return &user, nil

}
