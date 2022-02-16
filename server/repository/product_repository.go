package repository

import (
	"gorm.io/gorm"
	"server/entity"
	"server/exception"
)
type ProductRepository interface {
	CreateProduct(article *entity.Product, db *gorm.DB) (*entity.Product, *exception.ErrorResponse)
}