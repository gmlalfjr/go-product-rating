package repository

import (
	"server/collections"
	"server/entity"
	"server/exception"
)

type ProductRepositoryImpl struct {
	collections.Collection
}

func NewProductRepository(db collections.Collection) ProductRepository {
	return &ProductRepositoryImpl{
		db,
	}
}

func (p ProductRepositoryImpl) CreateProduct(article *entity.Product) (*entity.Product, *exception.ErrorResponse) {

	if err := p.ProductDB().Create(&article).Error; err !=nil {
		return nil, &exception.ErrorResponse{
			Code:   500,
			Status: "Internal Server Error",
			Data:   err.Error(),
		}
	}
	return article, nil
}
