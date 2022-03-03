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

func (p ProductServiceImpl) GetAllProduct() ([]product_web.ProductResponse, *exception.ErrorResponse) {
	data, err := p.productRepository.GetAllProduct()
	if err != nil {
		return nil, err
	}
	var result []product_web.ProductResponse
	for _, val := range data {
		result = append(result, product_web.ProductResponse{
			Title:       val.Title,
			Description: val.Description,
		})
	}

	return result, nil
}

func (p ProductServiceImpl) GetProductById(id int) (*product_web.ProductResponse, *exception.ErrorResponse) {
	data, err := p.productRepository.FindProductById(id)
	if err != nil {
		return nil, err
	}

	return 	&product_web.ProductResponse{
		Title:       data.Title,
		Description: data.Description,
	}, nil
}


