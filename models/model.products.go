package models

import (
	"time"
)

type ModelProduct struct {
	ID          string    `json:"id" gorm:"primary_key"`
	Name        string    `json:"name" gorm:"type:varchar;"`
	Price       string    `json:"price" gorm:"type:varchar;"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (ModelProduct) TableName() string {
	return "products"
}

// func (samethod *SegAnalisaMethod) FindAllSegAnalisaMethodsFull(db *gorm.DB, antipe uint32, search string, limit uint64, offset uint64) ([]SegAnalisaMethod, error) {
// 	search_str := "%" + search + "%"

// 	var err error
// 	dsamethod := []SegAnalisaMethod{}
// 	err = db.Debug().Model(&SegAnalisaMethod{}).Select("seg_analisa_methods.id, seg_analisa_methods.id_analisa_type, seg_analisa_methods.name, seg_analisa_methods.description, seg_analisa_methods.location, seg_analisa_methods.status_proyek_boost, b.location_name").Joins(" join seg_method_lokasi b on seg_analisa_methods.id = b.id").Where("id_analisa_type = ? and lower(name) like lower(?)", antipe, search_str).Order("id asc").Limit(limit).Offset(offset).Find(&dsamethod).Error
// 	if err != nil {
// 		return []SegAnalisaMethod{}, err
// 	}
// 	return dsamethod, err
// }
