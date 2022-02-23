package rating_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/exception"
	"server/services/rating_service"
	"server/web"
	"server/web/rating_web"
	"strconv"
)

type RatingControllerImpl struct {
	ratingService rating_service.RatingService
}

func NewRatingController(service rating_service.RatingService) RatingController {
	return RatingControllerImpl{ratingService: service}
}

func (r RatingControllerImpl) CreateRating(ctx *gin.Context) {
	productId := ctx.Param("productId")
	u64, errParse := strconv.ParseUint(productId, 10, 32)
	if errParse != nil {
		ctx.JSON(http.StatusBadGateway, exception.ErrorResponse{
			Code:   400,
			Status: "Body Request Error",
			Data:   errParse.Error(),
		})
		return
	}
	id := uint(u64)
	rating := &rating_web.RatingCreateRequest{
		ProductId: id,
	}

	if errBindJson := ctx.ShouldBindJSON(rating); errBindJson != nil {
		ctx.JSON(http.StatusBadGateway, exception.ErrorResponse{
			Code:   400,
			Status: "Body Request Error",
			Data:   errBindJson.Error(),
		})
		return
	}

	data, err := r.ratingService.CreateRating(rating)
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

func (r RatingControllerImpl) GetAllRatingByProduct(ctx *gin.Context) {
	productId := ctx.Param("productId")
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

	data, err := r.ratingService.GetAllProductRating(id)
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
