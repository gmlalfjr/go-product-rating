package rating_controller

import "github.com/gin-gonic/gin"

type RatingController interface {
	CreateRating(ctx *gin.Context)
	GetAllRatingByProduct(ctx *gin.Context)
}
