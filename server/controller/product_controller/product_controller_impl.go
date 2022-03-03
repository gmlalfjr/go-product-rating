package product_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/exception"
	"server/services/product_service"
	"server/web"
	"server/web/product_web"
	"strconv"
)

type ProductControllerImpl struct {
	productService product_service.ProductService
}

func NewProductController(service product_service.ProductService) ProductController {
	return ProductControllerImpl{productService: service}
}

func (p ProductControllerImpl) CreateProduct(ctx *gin.Context) {
	product := &product_web.ProductCreateRequest{}

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

func (p ProductControllerImpl) GetProduct(ctx *gin.Context) {
	data, err := p.productService.GetAllProduct()
	if err != nil {
		ctx.JSON(err.Code, exception.ErrorResponse{
			Code:   err.Code,
			Status: err.Status,
			Data:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, web.Response{
		Code:   http.StatusOK,
		Status: "Success Get Data",
		Data:   data,
	})
	return
}

func (p ProductControllerImpl) GetProductDetail(ctx *gin.Context) {
	productId := ctx.Param("id")
	u64, errParse := strconv.ParseUint(productId, 10, 32)
	if errParse != nil {
		ctx.JSON(http.StatusBadGateway, exception.ErrorResponse{
			Code:   400,
			Status: "Body Request Error",
			Data:   errParse.Error(),
		})
		return
	}
	id := int(u64)
	data, err := p.productService.GetProductById(id)
	if err != nil {
		ctx.JSON(err.Code, exception.ErrorResponse{
			Code:   err.Code,
			Status: err.Status,
			Data:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, web.Response{
		Code:   http.StatusOK,
		Status: "Success Get Data",
		Data:   data,
	})
	return
}

