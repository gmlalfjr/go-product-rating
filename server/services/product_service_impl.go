package services

import (
	"gorm.io/gorm"
	"server/entity"
	"server/exception"
	"server/repository"
	"server/web"
	"time"
)

type ProductServiceImpl struct {
	productRepository repository.ProductRepository
	DB *gorm.DB
}

func NewProductService(repo repository.ProductRepository, db *gorm.DB) ProductService {
	return ProductServiceImpl{
		productRepository: repo,
		DB:                db,
	}
}

func (p ProductServiceImpl) CreateProduct(request *web.ProductCreateRequest) (*web.ProductCreateResponse, *exception.ErrorResponse) {
	product := entity.Product{
		Title:       request.Title,
		Description: request.Description,
		CreatedAt:   time.Now(),
		ModifiedAt:  time.Now(),
	}
	data, err :=p.productRepository.CreateProduct(&product, p.DB)
	if err != nil {
		return nil, err
	}

	return &web.ProductCreateResponse{
		Title: data.Title,
		Description:  data.Description,
	}, nil
}
