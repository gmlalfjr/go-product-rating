package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/exception"
	"server/services"
	"server/web"
)

type ProductControllerImpl struct {
	productService services.ProductService
}

func NewProductController(service services.ProductService) ProductController {
	return ProductControllerImpl{productService: service}
}

func (p ProductControllerImpl) CreateProduct(ctx *gin.Context) {
	product := &web.ProductCreateRequest{}

	if errBindJson := ctx.ShouldBindJSON(product); errBindJson != nil {
		ctx.JSON(http.StatusBadGateway, exception.ErrorResponse{
			Code:   400,
			Status: "Body Request Error",
			Data:   errBindJson.Error(),
		})
		return
	}

	data, err := p.productService.CreateProduct(product)
	if err != nil {
		ctx.JSON(err.Code, exception.ErrorResponse{
			Code:   err.Code,
			Status: err.Status,
			Data:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, web.Response{
		Code:   http.StatusCreated,
		Status: "Success Create Data",
		Data:   data,
	})
	return
}
