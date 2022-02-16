package repository

import (
	"server/entity"
	"server/exception"
)
type ProductRepository interface {
	CreateProduct(article *entity.Product) (*entity.Product, *exception.ErrorResponse)
}