package repository

import (
	"github.com/amarmaulana95/go-echo/models"
	"github.com/amarmaulana95/go-echo/schemas"

	"gorm.io/gorm"
)

type repositoryProduct struct {
	db *gorm.DB
}

// EntityCreate implements entity.Entityproductss
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

func (r *repositoryProduct) EntityResultAll(search string, limit uint64, offset uint64) (*[]models.ModelProduct, error) {
	searchs := "%" + search + "%"
	aslimit := limit
	asoffset := offset
	var products []models.ModelProduct
	r.db.Debug().Model(&products).Select("id, name, price, description, created_at").Where("lower(name) like lower(?)", searchs).Order("id asc").Limit(int(aslimit)).Offset(int(asoffset)).Find(&products)
	return &products, nil
}

func (r *repositoryProduct) EntityResultAllTotal(search string) uint64 {
	cari := "%" + search + "%"
	tampung := []models.ModelProduct{}
	result := r.db.Debug().Where("lower(name) like lower(?)", cari).Find(&tampung)
	hasil := uint64(result.RowsAffected)
	return hasil
}
