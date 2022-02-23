package product_service

import (
	"server/entity"
	"server/exception"
	"server/repository/product_repository"
	"server/web/product_web"
	"time"
)

type ProductServiceImpl struct {
	productRepository product_repository.ProductRepository
}

func NewProductService(repo product_repository.ProductRepository) ProductService {
	return ProductServiceImpl{
		productRepository: repo,
	}
}

func (p ProductServiceImpl) CreateProduct(request *product_web.ProductCreateRequest) (*product_web.ProductCreateResponse, *exception.ErrorResponse) {
	product := entity.Product{
		Title:       request.Title,
		Description: request.Description,
		CreatedAt:   time.Now(),
		ModifiedAt:  time.Now(),
	}
	data, err := p.productRepository.CreateProduct(&product)
	if err != nil {
		return nil, err
	}

	return &product_web.ProductCreateResponse{
		Title:       data.Title,
		Description: data.Description,
	}, nil
}
