package repository

import (
	"gorm.io/gorm"
	"server/entity"
	"server/exception"
)

type ProductRepositoryImpl struct {

}

func NewProductRepository() ProductRepository {
	return &ProductRepositoryImpl{}
}

func (p ProductRepositoryImpl) CreateProduct(article *entity.Product, db *gorm.DB) (*entity.Product, *exception.ErrorResponse) {
	if err := db.Table("products").Create(&article).Error; err !=nil {
		return nil, &exception.ErrorResponse{
			Code:   500,
			Status: "Internal Server Error",
			Data:   err.Error(),
		}
	}
	return article, nil
}
