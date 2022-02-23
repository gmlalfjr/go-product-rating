package product_service

import (
	"server/exception"
	"server/web/product_web"
)

type ProductService interface {
	CreateProduct(request *product_web.ProductCreateRequest) (*product_web.ProductCreateResponse, *exception.ErrorResponse)
}
