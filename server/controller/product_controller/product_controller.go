package product_controller

import "github.com/gin-gonic/gin"

type ProductController interface {
	CreateProduct(ctx *gin.Context)
	GetProduct(ctx *gin.Context)
	GetProductDetail(ctx *gin.Context)
}
