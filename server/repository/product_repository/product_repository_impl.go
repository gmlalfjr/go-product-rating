package product_repository

import (
	"server/collections"
	"server/entity"
	"server/exception"
)

type ProductRepositoryImpl struct {
	collections.Collection
}

func NewProductRepository(db collections.Collection) ProductRepository {
	return &ProductRepositoryImpl{
		db,
	}
}

func (p ProductRepositoryImpl) CreateProduct(product *entity.Product) (*entity.Product, *exception.ErrorResponse) {

	if err := p.ProductDB().Create(&product).Error; err != nil {
		return nil, &exception.ErrorResponse{
			Code:   500,
			Status: "Internal Server Error",
			Data:   err.Error(),
		}
	}
	return product, nil
}

func (p ProductRepositoryImpl) FindProductById(id int) (*entity.Product, *exception.ErrorResponse) {
	var product *entity.Product
	if err := p.ProductDB().Where("id = ?", id).First(&product).Error; err != nil {
		return nil, &exception.ErrorResponse{
			Code:   500,
			Status: "Internal Server Error",
			Data:   err.Error(),
		}
	}
	return product, nil
}

func (p ProductRepositoryImpl) GetAllProduct() ([]entity.Product, *exception.ErrorResponse) {
	var product []entity.Product
	if err := p.ProductDB().Find(&product).Error; err !=nil {
		return nil, &exception.ErrorResponse{
			Code:   500,
			Status: "Internal Server Error",
			Data:   err.Error(),
		}
	}

	return product, nil
}


