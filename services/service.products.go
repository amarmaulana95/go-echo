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

func (s *serviceProduct) EntityResultAll(search string, limit uint64, offset uint64) (*[]models.ModelProduct, error) {
	res, err := s.products.EntityResultAll(search, limit, offset)
	return res, err
}

func (s *serviceProduct) EntityResultAllTotal(search string) uint64 {
	res := s.products.EntityResultAllTotal(search)
	return res
}
