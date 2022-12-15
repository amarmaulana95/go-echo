package repository

import (
	"github.com/amarmaulana95/go-echo/models"
	"github.com/amarmaulana95/go-echo/schemas"

	"gorm.io/gorm"
)

type repositoryProduct struct {
	db *gorm.DB
}

// EntityCreate implements entity.EntityUsers
func (*repositoryProduct) EntityCreate(input *schemas.SchemaProduct) (*models.ModelProduct, error) {
	panic("unimplemented")
}

func NewRepositoryProduct(db *gorm.DB) *repositoryProduct {
	return &repositoryProduct{
		db: db,
	}
}

func (r *repositoryProduct) EntityResults() (*[]models.ModelProduct, error) {
	var products []models.ModelProduct
	db := r.db.Model(&products)
	err := db.Debug().Find(&products).Error
	if err != nil {
		return &products, err
	}
	return &products, nil
}
