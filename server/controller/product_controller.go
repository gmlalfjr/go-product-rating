package controller

import "github.com/gin-gonic/gin"

type ProductController interface {
	CreateProduct(ctx *gin.Context)
}
