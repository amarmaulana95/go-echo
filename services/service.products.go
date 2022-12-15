package services

import (
	"github.com/amarmaulana95/go-echo/entity"
	"github.com/amarmaulana95/go-echo/models"
)

type serviceProduct struct {
	products entity.EntityProducts
}

func NewserviceProduct(products entity.EntityProducts) *serviceProduct {
	return &serviceProduct{
		products: products,
	}
}

func (s *serviceProduct) EntityResults() (*[]models.ModelProduct, error) {
	res, err := s.products.EntityResults()
	if err != nil {
		return res, err
	}
	return res, err
}
