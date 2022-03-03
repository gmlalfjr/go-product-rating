package product_service

import (
	"server/exception"
	"server/web/product_web"
)

type ProductService interface {
	CreateProduct(request *product_web.ProductCreateRequest) (*product_web.ProductCreateResponse, *exception.ErrorResponse)
	GetAllProduct() ([]product_web.ProductResponse, *exception.ErrorResponse)
	GetProductById(id int) (*product_web.ProductResponse, *exception.ErrorResponse)
}
