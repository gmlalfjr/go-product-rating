package services

import (
	"server/exception"
	"server/web"
)

type ProductService interface {
	CreateProduct(request *web.ProductCreateRequest)(*web.ProductCreateResponse, *exception.ErrorResponse)
}