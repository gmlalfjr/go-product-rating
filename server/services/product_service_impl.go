package services

import (
	"server/entity"
	"server/exception"
	"server/repository"
	"server/web"
	"time"
)

type ProductServiceImpl struct {
	productRepository repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return ProductServiceImpl{
		productRepository: repo,
	}
}

func (p ProductServiceImpl) CreateProduct(request *web.ProductCreateRequest) (*web.ProductCreateResponse, *exception.ErrorResponse) {
	product := entity.Product{
		Title:       request.Title,
		Description: request.Description,
		CreatedAt:   time.Now(),
		ModifiedAt:  time.Now(),
	}
	data, err :=p.productRepository.CreateProduct(&product)
	if err != nil {
		return nil, err
	}

	return &web.ProductCreateResponse{
		Title: data.Title,
		Description:  data.Description,
	}, nil
}
