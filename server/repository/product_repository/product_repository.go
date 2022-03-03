package product_repository

import (
	"server/entity"
	"server/exception"
)

type ProductRepository interface {
	CreateProduct(product *entity.Product) (*entity.Product, *exception.ErrorResponse)
	FindProductById(id int) (*entity.Product, *exception.ErrorResponse)
	GetAllProduct() ([]entity.Product, *exception.ErrorResponse)
}
