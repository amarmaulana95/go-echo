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
	if err != nil {
		return &users, err
	}
	return &users, nil
}

func (r *repositoryUser) EntityResultID(input *schemas.SchemaUser) (*models.ModelUser, error) {
	var user models.ModelUser
	user.ID = input.ID
	db := r.db.Model(&user)
	checkUser := db.Debug().First(&user, "id=?", input.ID)
	if checkUser.RowsAffected < 1 {
		return &user, checkUser.Error
	}

	return &user, nil
}

func (r *repositoryUser) EntityResultIDSearch(search string) (*[]models.ModelUser, error) {

	searchs := "%" + search + "%"
	// aslimit := limit
	// asoffset := offset

	var user []models.ModelUser

	r.db.Debug().Model(&user).Select("id, username, email").Where("lower(username) like lower(?)", searchs).Order("id asc").Find(&user)
	return &user, nil
}

func (r *repositoryUser) EntityResultAll(search string, limit uint64, offset uint64) (*[]models.ModelUser, error) {
	searchs := "%" + search + "%"
	aslimit := limit
	asoffset := offset
	var user []models.ModelUser
	r.db.Debug().Model(&user).Select("id, username, email").Where("lower(username) like lower(?)", searchs).Order("id asc").Limit(int(aslimit)).Offset(int(asoffset)).Find(&user)
	return &user, nil
}

// func (r *repositoryUser) EntityResultAllTotal(search string) (*[]models.ModelUser, error) {
// 	searchs := "%" + search + "%"
// 	var user []models.ModelUser
// 	r.db.Where("username = ?", searchs).Find(&user)
// 	return &user, nil
// }

// func (r *repositoryUser) EntityResultAllTotal(search string) uint64 {
// 	searchs := "%" + search + "%"
// 	var user []models.ModelUser
// 	result := r.db.Where("username = ?", searchs).Find(&user)
// 	// return &user, nil
// 	return uint64(result.RowsAffected)
// }
